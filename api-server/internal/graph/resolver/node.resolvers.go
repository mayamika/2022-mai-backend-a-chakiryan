package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return ent.FromContext(ctx).Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return ent.FromContext(ctx).Noders(ctx, ids)
}

// Query returns gqlgenerated.QueryResolver implementation.
func (r *Resolver) Query() gqlgenerated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
