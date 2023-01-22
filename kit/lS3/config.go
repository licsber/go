package lS3

import "os"

type Config struct {
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
	EnableSSL bool
}

func (c *Config) Must() {
	if c.Endpoint == "" {
		c.Endpoint = os.Getenv("L_S3_ENDPOINT")
		if c.Endpoint == "" {
			panic("lS3 need L_S3_ENDPOINT env var.")
		}
	}

	if c.Region == "" {
		c.Region = os.Getenv("L_S3_REGION")
		if c.Region == "" {
			panic("lS3 need L_S3_REGION env var.")
		}
	}

	if c.AccessKey == "" {
		c.AccessKey = os.Getenv("L_S3_ACCESS")
		if c.AccessKey == "" {
			panic("lS3 need L_S3_ACCESS env var.")
		}
	}

	if c.SecretKey == "" {
		c.SecretKey = os.Getenv("L_S3_SECRET")
		if c.SecretKey == "" {
			panic("lS3 need L_S3_SECRET env var.")
		}
	}
}
