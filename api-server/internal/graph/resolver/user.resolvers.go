package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
)

func (r *userResolver) Friends(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy ent.UserOrder) (*ent.UserConnection, error) {
	return ent.FromContext(ctx).User.QueryFriends(obj).
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(&orderBy),
		)
}

// User returns gqlgenerated.UserResolver implementation.
func (r *Resolver) User() gqlgenerated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
