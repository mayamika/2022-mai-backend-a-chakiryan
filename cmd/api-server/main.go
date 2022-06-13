package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := app.Config{
		Addr: ":8080",
	}
	config.BindEnv()

	logger := newLogger()

	startCtx, cancelStart := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelStart()

	a, err := app.New(startCtx, config, logger)
	if err != nil {
		cancelStart()
		logger.Fatal("start", zap.Error(err))
	}
	logger.Info("started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logger.Info("stopping")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Stop(ctx); err != nil {
		logger.Error("shutdown", zap.Error(err))
	}
}

func newLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Development = false
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
