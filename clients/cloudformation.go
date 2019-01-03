package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudformation"
    "github.com/aws/aws-sdk-go/aws/awserr"
)

type CfnClient struct {
    cli *cloudformation.CloudFormation
}

func NewCloudformation(sess *session.Session) *CfnClient {
    client := cloudformation.New(sess)

    return &CfnClient{ cli: client }
}

func (cfn *CfnClient) ListStacks(input *cloudformation.ListStacksInput) *cloudformation.ListStacksOutput {
    resp, err := cfn.cli.ListStacks(input)
    if err != nil {
       handleError(err)
    }
    return resp
}

func (cfn *CfnClient) ListStackReources(input *cloudformation.ListStackResourcesInput) *cloudformation.ListStackResourcesOutput {
    resp, err := cfn.cli.ListStackResources(input)
    if err != nil {
        handleError(err)
    }
    return resp
}

func handleError(err error) {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
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
