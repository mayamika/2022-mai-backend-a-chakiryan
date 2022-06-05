package gqlresolver

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/feed"
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

type FeedController interface {
	Feed(
		ctx context.Context,
		first int,
		after *string,
		search *string,
	) (*feed.FeedPayload, error)
	PublishPost(ctx context.Context, input feed.PostInput) (*feed.Post, error)
}

type Resolver struct {
	authController   AuthController
	friendController FriendController
	feedController   FeedController
}

func New(
	authController AuthController,
	friendController FriendController,
	feedController FeedController,
) *Resolver {
	return &Resolver{
		authController:   authController,
		friendController: friendController,
		feedController:   feedController,
	}
}
