package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	api := chi.NewRouter()
	api.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		type HelloResponse struct {
			Hello string `json:"hello,omitempty"`
		}
		res := &HelloResponse{
			Hello: "World",
		}

		render.JSON(w, r, res)
		render.Status(r, http.StatusOK)
	})

	router := chi.NewRouter()
	router.Mount("/api", api)

	corsOptions := cors.Options{}

	server := &http.Server{
		Addr:    ":8080",
		Handler: cors.New(corsOptions).Handler(router),
	}
	logger := newLogger()

	go func() {
		logger.Info("listening http", zap.String("addr", server.Addr))
		defer logger.Info("stopped")

		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("listen http", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logger.Info("stopping")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("shutdown http", zap.Error(err))
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
