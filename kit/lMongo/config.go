package lMongo

import (
	"context"
	"fmt"
	"os"
)

type Config struct {
	CTX context.Context
	URI string
}

func (c *Config) GetURI() string {
	if c.URI != "" {
		return c.URI
	}

	c.URI = os.Getenv("L_MONGO_URI")
	fmt.Println(c)
	fmt.Println(os.Environ())
	if c.URI == "" {
		panic("lMongo need L_MONGO_URI env var.")
	}

	return c.URI
}
