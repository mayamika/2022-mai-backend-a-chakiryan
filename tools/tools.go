//go:build tools

package tools

import (
	// GraphQL.
	_ "github.com/99designs/gqlgen"

	// Database.
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/volatiletech/null/v8"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)
