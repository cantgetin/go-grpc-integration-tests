package app

import (
	"context"
	"fmt"
	"go-integration-tests/internal/app/booksserviceapi"
	userserviceapi "go-integration-tests/internal/app/usersserviceapi"
	"go-integration-tests/internal/bootstrap"
	"go-integration-tests/internal/config"
	"go-integration-tests/internal/database/repository/bookrepository"
	"go-integration-tests/internal/database/repository/userrepository"
	"go-integration-tests/pkg/books"
	"go-integration-tests/pkg/users"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) error {

	// grpc server
	s := grpc.NewServer()
	_, cancel := context.WithCancel(context.Background())

	// db
	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to init db, %v", err)
	}

	l, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen tcp, %v", err)
	}

	// repos
	br := bookrepository.New(db)
	ur := userrepository.New(db)

	// services
	bs := booksserviceapi.New(br)
	us := userserviceapi.New(ur)

	books.RegisterBookServiceServer(s, bs)
	users.RegisterUserServiceServer(s, us)

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatalf("failed to service grpc server, %v", err)
		}
		log.Println("Server started")
	}()

	gracefulShutDown(s, cancel)

	return nil
}

func gracefulShutDown(s *grpc.Server, cancel context.CancelFunc) {
	const waitTime = 5 * time.Second

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("Received shutdown signal %v", sig)
	log.Println(errorMessage)
	s.GracefulStop()
	cancel()
	time.Sleep(waitTime)
}
