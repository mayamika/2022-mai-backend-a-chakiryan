package app

import (
	"os"
	"strings"
)

type Config struct {
	Addr                string
	Postgres            string
	OpensearchAddresses []string
	OpensearchUsername  string
	OpensearchPassword  string
}

func (c *Config) BindEnv() {
	c.Addr = os.Getenv("ADDR")
	c.Postgres = os.Getenv("POSTGRES")
	c.OpensearchAddresses = strings.Split(os.Getenv("OPENSEARCH_ADDRESSES"), ",")
	c.OpensearchUsername = os.Getenv("OPENSEARCH_USERNAME")
	c.OpensearchPassword = os.Getenv("OPENSEARCH_PASSWORD")
}
