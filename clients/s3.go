package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
    cli *s3.S3
}

func NewS3(sess *session.Session) *S3Client {
    client := s3.New(sess)

    return &S3Client{ cli: client }
}

func (s3cli *S3Client) ListBuckets(input *s3.ListBucketsInput) *s3.ListBucketsOutput {
    resp, err := s3cli.cli.ListBuckets(input)
    if err != nil {
        fmt.Println("Got error listing buckets:")
        fmt.Println(err.Error())
    }

    return resp
}

func (s3cli *S3Client) GetBucketPolicy(input *s3.GetBucketPolicyInput) *s3.GetBucketPolicyOutput {
    resp, err := s3cli.cli.GetBucketPolicy(input)
    if err != nil {
        fmt.Println("Got error listing buckets:")
        fmt.Println(err.Error())
    }

    return resp
}
