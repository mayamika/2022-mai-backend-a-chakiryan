package image

import (
	"context"

	"github.com/google/uuid"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/upload"
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	client *minio.Client
	bucket string
}

func NewStorage(client *minio.Client, bucket string) *Storage {
	return &Storage{
		client: client,
		bucket: bucket,
	}
}

func (s *Storage) AddImage(ctx context.Context, u upload.Upload) (string, error) {
	opts := minio.PutObjectOptions{
		ContentType: u.ContentType,
	}
	info, err := s.client.PutObject(ctx, s.bucket, s.generateName(), u.File, u.Size, opts)
	if err != nil {
		return "", err
	}
	return info.Key, nil
}

func (s *Storage) generateName() string {
	return uuid.NewString()
}
