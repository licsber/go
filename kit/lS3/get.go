package lS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func Get(disableSSL bool) *s3.S3 {
	endpoint := os.Getenv("L_S3_ENDPOINT")
	region := os.Getenv("L_S3_REGION")
	ak := os.Getenv("L_S3_ACCESS")
	sk := os.Getenv("L_S3_SECRET")

	cred := credentials.NewStaticCredentials(ak, sk, "")

	config := aws.NewConfig().
		WithEndpoint(endpoint).
		WithRegion(region).
		WithCredentials(cred).
		WithS3ForcePathStyle(true).
		WithDisableSSL(disableSSL)

	return s3.New(session.Must(session.NewSession(config)))
}
