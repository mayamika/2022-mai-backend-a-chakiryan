package middleware

import "net/http"

type MiddlewareFunc func(h http.Handler) http.Handler
