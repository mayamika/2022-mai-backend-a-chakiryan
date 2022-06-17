package app

import (
	"os"
	"strings"
)

type OpensearchConfig struct {
	Addresses []string
	Username  string
	Password  string
	CACert    string
}

type S3Config struct {
	Address         string
	AccessKeyID     string
	AccessKeySecret string
}

type Config struct {
	Port         string
	Postgres     string
	Opensearch   OpensearchConfig
	S3           S3Config
	ImagesBucket string
}

func (c *Config) BindEnv() {
	c.Port = os.Getenv("PORT")
	c.Postgres = os.Getenv("POSTGRES")
	c.Opensearch.Addresses = strings.Split(os.Getenv("OPENSEARCH_ADDRESSES"), ",")
	c.Opensearch.Username = os.Getenv("OPENSEARCH_USERNAME")
	c.Opensearch.Password = os.Getenv("OPENSEARCH_PASSWORD")
	c.Opensearch.CACert = os.Getenv("OPENSEARCH_CA_CERT")
	c.S3.Address = os.Getenv("S3_ADDRESS")
	c.S3.AccessKeyID = os.Getenv("S3_ACCESS_KEY_ID")
	c.S3.AccessKeySecret = os.Getenv("S3_ACCESS_KEY_SECRET")
	c.ImagesBucket = os.Getenv("IMAGES_BUCKET")
}
