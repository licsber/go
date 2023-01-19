package lS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Get(cfg *Config) *s3.S3 {
	cfg.Must()
	cred := credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, "")

	config := aws.NewConfig().
		WithEndpoint(cfg.Endpoint).
		WithRegion(cfg.Region).
		WithCredentials(cred).
		WithS3ForcePathStyle(true).
		WithDisableSSL(!cfg.EnableSSL)

	return s3.New(session.Must(session.NewSession(config)))
}
