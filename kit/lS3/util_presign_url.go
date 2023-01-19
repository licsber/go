package lS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

func PreSignUrl(s *s3.S3, bucket, key string, expires time.Duration) (string, error) {
	// 这里的返回值不是err 不要每次看过来都想改它（
	req, _ := s.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	return req.Presign(expires)
}
