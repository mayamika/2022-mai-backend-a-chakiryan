package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/feed"
)

func (r *queryResolver) Feed(ctx context.Context, after *string, first *int, before *string, last *int, search *string) (*feed.PostConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
