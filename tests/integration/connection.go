package integration

import (
	"context"
	"go-integration-tests/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func GetGrpcConnection(cfg *config.Config) (*grpc.ClientConn, error) {

	timeout := time.Second * 100
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"0.0.0.0"+cfg.GRPCAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}
