package gqlresolver

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

type AuthController interface {
	Login(context.Context, auth.LoginInput) (*auth.LoginPayload, error)
	Register(context.Context, auth.RegisterInput) (*auth.RegisterPayload, error)
}

type Resolver struct {
	authController AuthController
}

func New(authController AuthController) *Resolver {
	return &Resolver{
		authController: authController,
	}
}
