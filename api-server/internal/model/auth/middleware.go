package auth

import (
	"net/http"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

const authHeader = "Mai-Backend-Token"

func Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get(authHeader)

			v, err := token.FromSignedString(tokenStr)
			if err == nil {
				ctx := token.ToContext(r.Context(), v)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
