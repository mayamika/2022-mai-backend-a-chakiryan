package app

import (
	"context"
	"net/http"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/service/auth"
)

type App struct {
	httpServer *http.Server

	authService *auth.Service
}

func New(c Config) *App {
	a := &App{}

	r := a.routes()
	a.httpServer = &http.Server{
		Addr:    c.Addr,
		Handler: r,
	}

	return a
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
