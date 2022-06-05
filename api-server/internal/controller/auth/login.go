package auth

import (
	"context"
	"fmt"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/user"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
	"github.com/pkg/errors"
)

var ErrLoginFailed = errors.New("login attempt failed")

func (c *Controller) Login(ctx context.Context, input auth.LoginInput) (*auth.LoginPayload, error) {
	client := ent.FromContext(ctx)
	u, err := client.User.Query().
		Where(user.LoginEQ(input.Login)).
		Only(ctx)
	if err != nil {
		return nil, ErrLoginFailed
	}

	ok := c.compareHashAndPassword(u.PasswordHash, input.Password)
	if !ok {
		return nil, ErrLoginFailed
	}
	token, err := c.tokenFromUser(u)
	if err != nil {
		return nil, fmt.Errorf("generate auth token: %w", err)
	}

	return &auth.LoginPayload{
		Token: token,
	}, nil
}
