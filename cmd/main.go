package main

import (
	"context"
	"github.com/caarlos0/env"
	"go-integration-tests/internal/app"
	"go-integration-tests/internal/config"
	"log"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Fatalf("error running grpc server, %v", err)
	}
}
