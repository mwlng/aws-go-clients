package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSClient struct {
	cli *sqs.SQS
}

func NewSQS(sess *session.Session) *SQSClient {
	client := sqs.New(sess)

	return &SQSClient{cli: client}
}

func (sqsCli *SQSClient) CreateQueue() *sqs.CreateQueueOutput {
	input := &sqs.CreateQueueInput{}

	resp, err := sqsCli.cli.CreateQueue(input)
	if err != nil {
		sqsCli.handleError(err)
	}

	return resp
}

func (sqsCli *SQSClient) DeleteQueue() *sqs.DeleteQueueOutput {
	input := &sqs.DeleteQueueInput{}

	resp, err := sqsCli.cli.DeleteQueue(input)
	if err != nil {
		sqsCli.handleError(err)
	}

	return resp
}

func (sqsCli *SQSClient) ReceiveMessage() *sqs.ReceiveMessageOutput {
	input := &sqs.ReceiveMessageInput{}

	resp, err := sqsCli.cli.ReceiveMessage(input)
	if err != nil {
		sqsCli.handleError(err)
	}

	return resp
}

func (sqsCli *SQSClient) SendMessage() *sqs.SendMessageOutput {
	input := &sqs.SendMessageInput{}

	resp, err := sqsCli.cli.SendMessage(input)
	if err != nil {
		sqsCli.handleError(err)
	}

	return resp
}

func (sqsCli *SQSClient) SendMessageBatch() *sqs.SendMessageBatchOutput {
	input := &sqs.SendMessageBatchInput{}

	resp, err := sqsCli.cli.SendMessageBatch(input)
	if err != nil {
		sqsCli.handleError(err)
	}

	return resp
}

func (sqsCli *SQSClient) handleError(err error) {
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
}
