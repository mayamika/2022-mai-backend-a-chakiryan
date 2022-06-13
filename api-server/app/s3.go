package app

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func newS3Client(c S3Config) (*minio.Client, error) {
	return minio.New(c.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKeyID, c.AccessKeySecret, ""),
		Secure: true,
	})
}
