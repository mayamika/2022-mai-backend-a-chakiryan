package friend

import (
	"context"
	"errors"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/user"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) AddFriendRequest(ctx context.Context, to int) (*ent.FriendRequest, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, auth.ErrNotAuthenticated
	}

	if to == t.UserID {
		return nil, errors.New("can't add yourserlf to friends")
	}

	return ent.FromContext(ctx).FriendRequest.Create().
		SetFromID(t.UserID).
		SetToID(to).
		Save(ctx)
}

func (c *Controller) AcceptFriendRequest(ctx context.Context, id int) (int, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return 0, auth.ErrNotAuthenticated
	}

	client := ent.FromContext(ctx)
	r, err := client.FriendRequest.Get(ctx, id)
	if err != nil {
		return 0, err
	}
	from, err := r.From(ctx)
	if err != nil {
		return 0, err
	}
	to, err := r.To(ctx)
	if err != nil {
		return 0, err
	}

	if to.ID != t.UserID {
		return 0, auth.ErrPermissionDenied
	}

	if err := from.Update().AddFriends(to).Exec(ctx); err != nil {
		return 0, err
	}

	return r.ID, client.FriendRequest.DeleteOne(r).Exec(ctx)
}

func (c *Controller) DeclineFriendRequest(ctx context.Context, id int) (int, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return 0, auth.ErrNotAuthenticated
	}

	client := ent.FromContext(ctx)
	r, err := client.FriendRequest.Get(ctx, id)
	if err != nil {
		return 0, err
	}
	to, err := r.To(ctx)
	if err != nil {
		return 0, err
	}

	if to.ID != t.UserID {
		return 0, auth.ErrPermissionDenied
	}

	return id, client.FriendRequest.DeleteOne(r).Exec(ctx)
}

func (c *Controller) RemoveFriend(ctx context.Context, id int) (*ent.User, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, auth.ErrNotAuthenticated
	}

	client := ent.FromContext(ctx)
	err := client.User.Update().
		Where(user.ID(t.UserID)).
		RemoveFriendIDs(id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return client.User.Get(ctx, id)
}
