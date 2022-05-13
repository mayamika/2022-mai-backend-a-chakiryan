package app

import "os"

type Config struct {
	Addr string
}

func (c *Config) BindEnv() {
	c.Addr = os.Getenv("ADDR")
}
