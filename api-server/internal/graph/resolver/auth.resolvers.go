package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

func (r *mutationResolver) Login(ctx context.Context, input auth.LoginInput) (*auth.LoginPayload, error) {
	return r.authController.Login(ctx, input)
}

func (r *mutationResolver) Register(ctx context.Context, input auth.RegisterInput) (*auth.RegisterPayload, error) {
	return r.authController.Register(ctx, input)
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	return r.authController.Me(ctx)
}

// Mutation returns gqlgenerated.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgenerated.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgenerated.QueryResolver implementation.
func (r *Resolver) Query() gqlgenerated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
