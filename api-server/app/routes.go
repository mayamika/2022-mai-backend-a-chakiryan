package app

import (
	"context"
	"sync"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	gqlresolver "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/resolver"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

func (a *App) routes() chi.Router {
	resolver := gqlresolver.New(
		a.authController,
		a.friendController,
		a.feedController,
	)
	query := handler.NewDefaultServer(gqlgenerated.NewExecutableSchema(
		gqlgenerated.Config{
			Resolvers: resolver,
		},
	))
	query.Use(transactioner{
		client:        a.client,
		Transactioner: entgql.Transactioner{TxOpener: a.client},
	})

	r := chi.NewRouter()
	r.Use(auth.Middleware())

	r.Mount(`/query`, query)
	r.Mount(`/playground`, playground.Handler("GraphQL playground", `/query`))
	return r
}

var _ interface {
	graphql.HandlerExtension
	graphql.OperationContextMutator
	graphql.ResponseInterceptor
} = transactioner{}

type transactioner struct {
	client *ent.Client
	entgql.Transactioner
}

// MutateOperationContext serializes field resolvers during mutations and queries.
func (transactioner) MutateOperationContext(_ context.Context, oc *graphql.OperationContext) *gqlerror.Error {
	if op := oc.Operation; op != nil && op.Operation == ast.Mutation || op.Operation == ast.Query {
		previous := oc.ResolverMiddleware
		var mu sync.Mutex
		oc.ResolverMiddleware = func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
			mu.Lock()
			defer mu.Unlock()
			return previous(ctx, next)
		}
	}
	return nil
}

// InterceptResponse puts regular client if operation type is Query or calls ent.Transactioner.
func (t transactioner) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	if op := graphql.GetOperationContext(ctx).Operation; op == nil || op.Operation != ast.Query {
		return t.Transactioner.InterceptResponse(ctx, next)
	}
	return next(ent.NewContext(ctx, t.client))
}
