package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

func (r *mutationResolver) Login(ctx context.Context, input auth.LoginInput) (*auth.LoginPayload, error) {
	return r.authService.Login(ctx, input)
}

func (r *mutationResolver) Register(ctx context.Context, input auth.RegisterInput) (*auth.RegisterPayload, error) {
	return r.authService.Register(ctx, input)
}

// Mutation returns gqlgenerated.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgenerated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
