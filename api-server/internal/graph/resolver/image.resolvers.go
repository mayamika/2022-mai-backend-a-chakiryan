package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	gqlmodel "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/model"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/upload"
)

func (r *mutationResolver) UploadImage(ctx context.Context, file graphql.Upload) (*gqlmodel.Image, error) {
	name, err := r.imageStorage.AddImage(ctx, upload.Upload(file))
	if err != nil {
		return nil, err
	}
	return &gqlmodel.Image{
		Name: name,
	}, nil
}
