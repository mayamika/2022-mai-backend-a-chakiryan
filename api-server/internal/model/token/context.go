package token

import "context"

type ctxMarker struct{}

var ctxMarkerKey = &ctxMarker{}

func ToContext(ctx context.Context, t Token) context.Context {
	return context.WithValue(ctx, ctxMarkerKey, t)
}

func FromContext(ctx context.Context) (Token, bool) {
	v := ctx.Value(ctxMarkerKey)
	if v == nil {
		return Token{}, false
	}
	t, ok := v.(Token)
	return t, ok
}
