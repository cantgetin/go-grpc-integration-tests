package tests

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"go-integration-tests/internal/app"
	"go-integration-tests/internal/config"
	"go-integration-tests/tests/integration"
	"google.golang.org/grpc"
	"testing"
)

func TestIntegration(t *testing.T) {
	suite.Run(t, &IntegrationSuite{})
}

type IntegrationSuite struct {
	suite.Suite
	cfg       *config.Config
	ctx       context.Context
	cancelCtx func()
	grpcConn  *grpc.ClientConn
	closers   []func() error
}

func (s *IntegrationSuite) SetupSuite() {
	s.cfg = &config.Config{}
	if err := env.Parse(s.cfg); err != nil {
		s.T().Fatalf("failed to parse cfg, %v", err)
	}

	s.ctx, s.cancelCtx = context.WithCancel(context.Background())

	s.T().Log("Starting Docker containers...")
	pool, dockerClose := integration.Start(s.T(), s.cfg)
	s.closers = append(s.closers, dockerClose)

	s.T().Log("Initializing DB with migrations...")
	closeDB := integration.InitDB(s.T(), pool, s.cfg)
	s.closers = append(s.closers, closeDB)

	s.T().Log("Setup completed")

	go func() {
		if err := app.Run(s.ctx, s.cfg); err != nil {
			s.T().Logf("application has exited %v", err)
		}
	}()

	conn, err := integration.GetGrpcConnection(s.cfg)
	if err != nil {
		s.T().Fatalf("Error getting grpc connection...")
	}
	s.grpcConn = conn
	s.closers = append(s.closers, conn.Close)

}

func (s *IntegrationSuite) TearDownSuite() {
	s.T().Log("Suite teardown...")
	s.cancelCtx()
	s.closeAll()
}

func (s *IntegrationSuite) closeAll() {
	for _, c := range s.closers {
		if err := c(); err != nil {
			s.Error(err)
		}
	}
}
