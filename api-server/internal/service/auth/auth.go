package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/user"
)

type TokenManager interface {
	UserFromToken(token string) (*user.User, error)
	CreateToken(userID string) (string, error)
}

type UserManager interface {
	User(ctx context.Context, id string) (*user.User, error)
	UserByLogin(ctx context.Context, login string) (*user.User, error)
}

type Service struct {
	userManager UserManager
}

func NewService(userManager UserManager) *Service {
	return &Service{
		userManager: userManager,
	}
}

func (s *Service) Login(ctx context.Context, input auth.LoginInput) (*auth.LoginPayload, error) {
	u, err := s.userManager.UserByLogin(ctx, input.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}

	return &auth.LoginPayload{}, nil
}

func (s *Service) Register(ctx context.Context, input auth.RegisterInput) (*auth.RegisterPayload, error) {
	return nil, errors.New("wad")
}
