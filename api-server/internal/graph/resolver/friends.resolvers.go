package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/friendrequest"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/user"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

func (r *mutationResolver) AddFriendRequest(ctx context.Context, to int) (*ent.FriendRequest, error) {
	return r.friendController.AddFriendRequest(ctx, to)
}

func (r *mutationResolver) AcceptFriendRequest(ctx context.Context, id int) (int, error) {
	return r.friendController.AcceptFriendRequest(ctx, id)
}

func (r *mutationResolver) DeclineFriendRequest(ctx context.Context, id int) (int, error) {
	return r.friendController.DeclineFriendRequest(ctx, id)
}

func (r *mutationResolver) RemoveFriend(ctx context.Context, id int) (*ent.User, error) {
	return r.friendController.RemoveFriend(ctx, id)
}

func (r *queryResolver) FriendRequests(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.FriendRequestConnection, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, auth.ErrNotAuthenticated
	}
	return ent.FromContext(ctx).FriendRequest.Query().
		Where(friendrequest.HasToWith(user.ID(t.UserID))).
		Paginate(ctx, after, first, before, last)
}
