package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/controller/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/controller/friend"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/migrate"
)

type App struct {
	client     *ent.Client
	httpServer *http.Server

	authController   *auth.Controller
	friendController *friend.Controller
}

func New(ctx context.Context, c Config, logger *zap.Logger) (*App, error) {
	client, err := ent.Open(dialect.Postgres, c.Postgres)
	if err != nil {
		return nil, fmt.Errorf("open ent client: %w", err)
	}
	migrateOpts := []schema.MigrateOption{
		migrate.WithGlobalUniqueID(true),
	}
	if err := client.Schema.Create(ctx, migrateOpts...); err != nil {
		client.Close()
		return nil, fmt.Errorf("migrate schema: %w", err)
	}

	a := &App{
		client:           client,
		authController:   auth.NewController(),
		friendController: friend.NewController(),
	}

	r := a.routes()
	a.httpServer = &http.Server{
		Addr:    c.Addr,
		Handler: r,
	}
	go func() {
		err := a.httpServer.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("serve http", zap.Error(err))
		}
	}()

	return a, nil
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return a.client.Close()
}
