package token_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

func TestToken(t *testing.T) {
	want := token.Token{
		UserID: 124,
	}

	s, err := want.SignedString()
	require.NoError(t, err)

	actual, err := token.FromSignedString(s)
	require.NoError(t, err)

	require.Equal(t, want, actual)
}
