package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Neftik/project3/sso/internal/app/grpc"
	"github.com/Neftik/project3/sso/internal/service/auth"
	"github.com/Neftik/project3/sso/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, databaseUrl string, tokenTTL time.Duration) *App {
	storage, err := postgres.Connect(databaseUrl)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}
}
