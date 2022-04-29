package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/user"
)

type UserManager interface {
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

func (s *Service) Login(ctx context.Context, input *user.LoginInput) (*user.LoginPayload, error) {
	u, err := s.userManager.UserByLogin(ctx, input.Login)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}

	return &user.LoginPayload{}, nil
}

func (s *Service) UserFromToken(token string) (*user.User, error) {

}
