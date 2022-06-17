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
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/controller/feed"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/controller/friend"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/migrate"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/storage/image"
)

type App struct {
	client     *ent.Client
	httpServer *http.Server

	authController   *auth.Controller
	friendController *friend.Controller
	feedController   *feed.Controller
	imageStorage     *image.Storage
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

	opensearchClient, err := newOpensearchClient(c.Opensearch)
	if err != nil {
		return nil, fmt.Errorf("open opensearch client: %w", err)
	}
	if err := pingOpenSearch(ctx, opensearchClient); err != nil {
		logger.Warn("ping opensearch cluster failed", zap.Error(err))
	}

	feedController, err := feed.NewController(ctx, opensearchClient)
	if err != nil {
		return nil, fmt.Errorf("open feed controller: %w", err)
	}

	s3Client, err := newS3Client(c.S3)
	if err != nil {
		return nil, fmt.Errorf("open s3 client: %w", err)
	}

	a := &App{
		client:           client,
		authController:   auth.NewController(),
		friendController: friend.NewController(),
		feedController:   feedController,
		imageStorage:     image.NewStorage(s3Client, c.ImagesBucket),
	}

	a.httpServer = &http.Server{
		Addr:    ":" + c.Port,
		Handler: a.routes(),
	}
	go func() {
		err := a.httpServer.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Panic("serve http", zap.Error(err))
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
