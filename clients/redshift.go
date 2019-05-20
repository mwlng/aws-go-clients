package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
)

type RedShiftClient struct {
	cli *redshift.Redshift
}

func NewRedShift(sess *session.Session) *RedShiftClient {
	client := redshift.New(sess)

	return &RedShiftClient{cli: client}
}

func (rsCli *RedShiftClient) GetClusterCreds(clusterId *string,
	dbUser *string,
	dbGroup *[]*string,
	dbName *string) *redshift.GetClusterCredentialsOutput {
	input := &redshift.GetClusterCredentialsInput{}
	resp, err := rsCli.cli.GetClusterCredentials(input)
	if err != nil {
		rsCli.handleError(err)
		return nil
	}
	return resp
}

func (rsCli *RedShiftClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case redshift.ErrCodeClusterNotFoundFault:
			fmt.Println(redshift.ErrCodeClusterNotFoundFault, aerr.Error())
		case redshift.ErrCodeUnsupportedOperationFault:
			fmt.Println(redshift.ErrCodeClusterNotFoundFault, aerr.Error())
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
