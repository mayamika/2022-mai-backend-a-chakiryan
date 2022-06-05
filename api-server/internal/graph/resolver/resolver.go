package gqlresolver

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

type FriendController interface {
	AddFriendRequest(ctx context.Context, to int) (*ent.FriendRequest, error)
	AcceptFriendRequest(ctx context.Context, id int) (int, error)
	DeclineFriendRequest(ctx context.Context, id int) (int, error)
	RemoveFriend(ctx context.Context, id int) (*ent.User, error)
}

type AuthController interface {
	Me(context.Context) (*ent.User, error)
	Login(context.Context, auth.LoginInput) (*auth.LoginPayload, error)
	Register(context.Context, auth.RegisterInput) (*auth.RegisterPayload, error)
}

type Resolver struct {
	authController   AuthController
	friendController FriendController
}

func New(
	authController AuthController,
	friendController FriendController,
) *Resolver {
	return &Resolver{
		authController:   authController,
		friendController: friendController,
	}
}
