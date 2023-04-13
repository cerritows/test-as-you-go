package levelthree

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucketName = "bucket123"

type FileClient interface {
	GetObject(ctx context.Context, input *s3.GetObjectInput) (s3.GetObjectOutput, error)
}

func GetFile(client FileClient, filename string) error {

	output, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Key:    aws.String(filename),
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	doSomethingWithFile(output.Body)

	return nil
}

func doSomethingWithFile(io.ReadCloser) {}

// // Load the Shared AWS Configuration (~/.aws/config)
// cfg, err := config.LoadDefaultConfig(context.TODO())
// if err != nil {
// 	return err
// }

// // Create an Amazon S3 service client
// client := s3.NewFromConfig(cfg)

// // Get the first page of results for ListObjectsV2 for a bucket
// output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
// 	Key:    aws.String(filename),
// 	Bucket: aws.String(bucketName),
// })
// if err != nil {
// 	return err
// }
