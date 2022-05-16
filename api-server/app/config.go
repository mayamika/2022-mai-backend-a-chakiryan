package app

import "os"

type Config struct {
	Addr     string
	Postgres string
}

func (c *Config) BindEnv() {
	c.Addr = os.Getenv("ADDR")
	c.Postgres = os.Getenv("POSTGRES")
}
