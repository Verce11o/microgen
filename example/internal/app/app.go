package app

import (
	"context"
	"fmt"
	"github.com/Verce11o/example/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	cfg        *config.Config
	log        *zap.SugaredLogger
	grpcServer *grpc.Server
}

func New(ctx context.Context, cfg *config.Config, log *zap.SugaredLogger) *App {

	repo := repository.NewRepository(db, cfg)

	services := service.NewService(log, repo)

	server := grpc.NewServer()

	example.Register(log, services, server, trace.Tracer)

	return &App{
		cfg:        cfg,
		log:        log,
		grpcServer: server,
	}
}
