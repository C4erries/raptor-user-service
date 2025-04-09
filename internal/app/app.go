package app

import (
	"log/slog"

	grpcapp "github.com/c4erries/raptor-user-service/internal/app/gprc"
	"github.com/c4erries/raptor-user-service/internal/config"
	"github.com/c4erries/raptor-user-service/internal/database/postgres"
	"github.com/c4erries/raptor-user-service/internal/service/userinfo"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	StorageConfig config.StorageConfig,
) *App {
	storage, err := postgres.New(StorageConfig)
	if err != nil {
		panic(err)
	}

	service := userinfo.New(storage)

	grpcApp := grpcapp.New(log, service, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
