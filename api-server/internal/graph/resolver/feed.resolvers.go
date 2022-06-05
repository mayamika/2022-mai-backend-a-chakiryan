package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/feed"
)

func (r *mutationResolver) PublishPost(ctx context.Context, input feed.PostInput) (*feed.Post, error) {
	return r.feedController.PublishPost(ctx, input)
}

func (r *postResolver) From(ctx context.Context, obj *feed.Post) (*ent.User, error) {
	return ent.FromContext(ctx).User.Get(ctx, obj.From)
}

func (r *queryResolver) Feed(ctx context.Context, first int, after *string, search *string) (*feed.FeedPayload, error) {
	return r.feedController.Feed(ctx, first, after, search)
}

// Post returns gqlgenerated.PostResolver implementation.
func (r *Resolver) Post() gqlgenerated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
