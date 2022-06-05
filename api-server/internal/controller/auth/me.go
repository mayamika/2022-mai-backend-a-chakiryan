package auth

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

func (c *Controller) Me(ctx context.Context) (*ent.User, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, auth.ErrNotAuthenticated
	}
	u, err := ent.FromContext(ctx).User.Get(ctx, t.UserID)
	if err != nil {
		return nil, auth.ErrNotAuthenticated
	}
	return u, nil
}
