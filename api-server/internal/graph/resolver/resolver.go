package gqlresolver

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
)

type AuthService interface {
	Login(context.Context, auth.LoginInput) (*auth.LoginPayload, error)
	Register(context.Context, auth.RegisterInput) (*auth.RegisterPayload, error)
}

type Resolver struct {
	authService AuthService
}

func New(authService AuthService) *Resolver {
	return &Resolver{
		authService: authService,
	}
}
