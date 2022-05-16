//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	gql, err := entgql.NewExtension(
		entgql.WithConfigPath("../graph/gqlgen.yml"),
		entgql.WithSchemaPath("../graph/schema/ent.graphqls"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(gql),
		entc.FeatureNames("privacy"),
	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
