package auth

import "errors"

var (
	ErrNotAuthenticated = errors.New("not authenticated")
	ErrPermissionDenied = errors.New("permission denied")
)
