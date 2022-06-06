package app

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"go.uber.org/zap"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/controller/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/controller/feed"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/controller/friend"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/migrate"
)

type App struct {
	client     *ent.Client
	httpServer *http.Server

	authController   *auth.Controller
	friendController *friend.Controller
	feedController   *feed.Controller
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

	opensearchClient, err := newOpensearchClient(c)
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

	a := &App{
		client:           client,
		authController:   auth.NewController(),
		friendController: friend.NewController(),
		feedController:   feedController,
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

func newOpensearchClient(c Config) (*opensearch.Client, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = tlsConfig

	return opensearch.NewClient(opensearch.Config{
		Addresses: c.OpensearchAddresses,
		Username:  c.OpensearchUsername,
		Password:  c.OpensearchPassword,
		Transport: transport,
	})
}

func pingOpenSearch(ctx context.Context, client *opensearch.Client) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pingRequest := opensearchapi.PingRequest{}
	_, err := pingRequest.Do(ctx, client)
	return err
}
