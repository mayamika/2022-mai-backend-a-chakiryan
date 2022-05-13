package middleware

import (
	"context"
	"net/http"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/user"
)

const AuthTokenKey = "Auth-Token"

type AuthService interface {
	UserFromToken(ctx context.Context, token string) (*user.User, error)
}

func NewAuth(authService AuthService) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			t := r.Header.Get(AuthTokenKey)

			u, err := authService.UserFromToken(ctx, t)
			if err == nil {
				user.ToContext(ctx, u)
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
