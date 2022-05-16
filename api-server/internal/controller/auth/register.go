package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/user"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

var ErrUserAlreadyExist = errors.New("user already exist")

func (c *Controller) Register(ctx context.Context, input auth.RegisterInput) (*auth.RegisterPayload, error) {
	client := ent.FromContext(ctx)
	userExist, err := client.User.Query().Where(
		user.LoginEQ(input.Login),
	).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if userExist {
		return nil, ErrUserAlreadyExist
	}

	passwordHash, err := c.passwordHash(input.Password)
	if err != nil {
		return nil, fmt.Errorf("can't generate password hash: %w", err)
	}

	u, err := client.User.Create().
		SetLogin(input.Login).
		SetPasswordHash(passwordHash).
		SetName(input.Name).
		SetSurname(input.Surname).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't save user: %w", err)
	}

	token, err := c.tokenFromUser(u)
	if err != nil {
		return nil, fmt.Errorf("can't generate auth token: %w", err)
	}

	return &auth.RegisterPayload{
		Token: token,
	}, nil
}
