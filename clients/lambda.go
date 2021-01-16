package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaClient struct {
	cli *lambda.Lambda
}

func NewLambda(sess *session.Session) *LambdaClient {
	client := lambda.New(sess)

	return &LambdaClient{cli: client}
}

func (lambdaCli *LambdaClient) Invoke(functionName string, payload []byte, invocationType string) *int64 {
	input := &lambda.InvokeInput{
		FunctionName:   aws.String(functionName),
		InvocationType: aws.String(invocationType),
		Payload:        payload,
	}

	output, err := lambdaCli.cli.Invoke(input)
	if err != nil {
		lambdaCli.handleError(err)
	}

	return output.StatusCode
}

func (lambdaCli *LambdaClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case lambda.ErrCodeInvalidParameterValueException:
			fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
		case lambda.ErrCodeResourceConflictException:
			fmt.Println(lambda.ErrCodeResourceConflictException, aerr.Error())
		case lambda.ErrCodeResourceNotFoundException:
			fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
		case lambda.ErrCodeTooManyRequestsException:
			fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
		case lambda.ErrCodeServiceException:
			fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
