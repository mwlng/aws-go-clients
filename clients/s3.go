package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	cli *s3.S3
}

func NewS3(sess *session.Session) *S3Client {
	client := s3.New(sess)

	return &S3Client{cli: client}
}

func (s3cli *S3Client) ListBuckets(input *s3.ListBucketsInput) *s3.ListBucketsOutput {
	resp, err := s3cli.cli.ListBuckets(input)
	if err != nil {
		s3cli.handleError(err)
	}

	return resp
}

func (s3cli *S3Client) GetBucketPolicy(input *s3.GetBucketPolicyInput) *s3.GetBucketPolicyOutput {
	resp, err := s3cli.cli.GetBucketPolicy(input)
	if err != nil {
		s3cli.handleError(err)
	}

	return resp
}

func (s3cli *S3Client) HeadObject(bucket *string, key *string) *s3.HeadObjectOutput {
	input := &s3.HeadObjectInput{
		Bucket: bucket,
		Key:    key,
	}
	resp, err := s3cli.cli.HeadObject(input)
	if err != nil {
		s3cli.handleError(err)
		return nil
	}
	return resp
}

func (s3Cli *S3Client) ListObjects(bucket *string, pathPrefix *string, continuationToken *string) (*string, []*s3.Object) {
	var input *s3.ListObjectsV2Input
	if continuationToken == nil {
		input = &s3.ListObjectsV2Input{
			Bucket: bucket,
			Prefix: pathPrefix,
		}
	} else {
		input = &s3.ListObjectsV2Input{
			Bucket:            bucket,
			Prefix:            pathPrefix,
			ContinuationToken: continuationToken,
		}
	}

	resp, err := s3Cli.cli.ListObjectsV2(input)
	if err != nil {
		s3Cli.handleError(err)
		return nil, nil
	}

	var nextToken *string = nil
	if *resp.IsTruncated {
		nextToken = resp.NextContinuationToken
	}
	return nextToken, resp.Contents
}

func (s3Cli *S3Client) GetObjectAcl(bucket *string, key *string) *s3.GetObjectAclOutput {
	input := &s3.GetObjectAclInput{
		Bucket: bucket,
		Key:    key,
	}

	resp, err := s3Cli.cli.GetObjectAcl(input)
	if err != nil {
		s3Cli.handleError(err)
	}

	return resp
}

func (s3Cli *S3Client) PutObjectAcl(bucket *string, key *string, acl *string) {
	input := &s3.PutObjectAclInput{
		Bucket: bucket,
		Key:    key,
		ACL:    acl,
	}
	_, err := s3Cli.cli.PutObjectAcl(input)
	if err != nil {
		s3Cli.handleError(err)
	}
	return
}

func (s3Cli *S3Client) CopyObject(srcBucket *string, tgtBucket *string,
	srcKey *string, tgtKey *string) {
	input := &s3.CopyObjectInput{
		ACL:        aws.String("bucket-owner-full-control"),
		CopySource: aws.String(fmt.Sprintf("/%s/%s", *srcBucket, *srcKey)),
		Bucket:     tgtBucket,
		Key:        tgtKey,
	}

	_, err := s3Cli.cli.CopyObject(input)
	if err != nil {
		s3Cli.handleError(err)
	}
	return
}

func (s3Cli *S3Client) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case s3.ErrCodeNoSuchBucket:
			fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
	return
}
