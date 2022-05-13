package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/user"
)

func (r *userResolver) Username(ctx context.Context, obj *user.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns gqlgenerated.UserResolver implementation.
func (r *Resolver) User() gqlgenerated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
