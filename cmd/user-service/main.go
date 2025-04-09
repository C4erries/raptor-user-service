package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/c4erries/raptor-user-service/internal/app"
	"github.com/c4erries/raptor-user-service/internal/config"
)

var (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)

	log.Info("starting application", slog.Any("config", config))
	port, err := strconv.Atoi(config.GRPC.Port)
	if err != nil {
		log.Error("error while parsing port", slog.Any("error", err))
	}
	application := app.New(log, port, config.Storage)
	application.GRPCSrv.MustRun()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return log
}
