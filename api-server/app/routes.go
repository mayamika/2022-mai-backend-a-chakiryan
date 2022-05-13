package app

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	gqlresolver "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/resolver"
)

func (a *App) routes() chi.Router {
	resolver := gqlresolver.New(
		a.authService,
	)
	query := handler.NewDefaultServer(gqlgenerated.NewExecutableSchema(
		gqlgenerated.Config{
			Resolvers: resolver,
		},
	))

	r := chi.NewRouter()
	r.Mount(`/query`, query)
	r.Mount(`/playground`, playground.Handler("GraphQL playground", `/query`))
	return r
}
