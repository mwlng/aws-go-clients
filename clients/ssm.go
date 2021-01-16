package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type SSMClient struct {
	cli *ssm.SSM
}

func NewSSM(sess *session.Session) *SSMClient {
	client := ssm.New(sess)

	return &SSMClient{cli: client}
}

func (ssmCli *SSMClient) GetParameter(name string) string {
	input := &ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: aws.Bool(true),
	}

	resp, err := ssmCli.cli.GetParameter(input)
	if err != nil {
		ssmCli.handleError(err)

		return ""
	}

	return *resp.Parameter.Value
}

func (ssmCli *SSMClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case ssm.ErrCodeInternalServerError:
			fmt.Println(ssm.ErrCodeInternalServerError, aerr.Error())
		case ssm.ErrCodeInvalidKeyId:
			fmt.Println(ssm.ErrCodeInvalidKeyId, aerr.Error())
		case ssm.ErrCodeParameterNotFound:
			fmt.Println(ssm.ErrCodeParameterNotFound, aerr.Error())
		case ssm.ErrCodeParameterVersionNotFound:
			fmt.Println(ssm.ErrCodeParameterVersionNotFound, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}
