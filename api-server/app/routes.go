package app

import (
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	gqlresolver "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/resolver"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

func (a *App) routes() chi.Router {
	resolver := gqlresolver.New(
		a.authController,
	)
	query := handler.NewDefaultServer(gqlgenerated.NewExecutableSchema(
		gqlgenerated.Config{
			Resolvers: resolver,
		},
	))
	query.Use(entgql.Transactioner{TxOpener: a.client})

	r := chi.NewRouter()
	r.Use(auth.Middleware())

	r.Mount(`/query`, query)
	r.Mount(`/playground`, playground.Handler("GraphQL playground", `/query`))
	return r
}
