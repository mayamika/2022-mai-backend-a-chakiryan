//go:build tools

package tools

import (
	// Ent.
	_ "entgo.io/contrib/entgql"
	_ "entgo.io/ent/cmd/ent"

	// GraphQL.
	_ "github.com/99designs/gqlgen"
)
