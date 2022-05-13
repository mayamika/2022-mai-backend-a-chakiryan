package user

import "context"

type ctxMarker struct{}

var ctxMarkerKey = &ctxMarker{}

func ToContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, ctxMarkerKey, u)
}

func FromContext(ctx context.Context) *User {
	v := ctx.Value(ctxMarkerKey)
	u, ok := v.(*User)
	if !ok {
		return Anonymous
	}
	return u
}
