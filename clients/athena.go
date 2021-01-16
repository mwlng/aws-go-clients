package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

type AthenaClient struct {
	cli *athena.Athena
}

func NewAthena(sess *session.Session) *AthenaClient {
	client := athena.New(sess)

	return &AthenaClient{cli: client}
}

func (athenaCli *AthenaClient) StartQueryExecution(catalogDB, query *string) (*string, error) {
	input := &athena.StartQueryExecutionInput{
		QueryExecutionContext: &athena.QueryExecutionContext{
			Database: catalogDB,
		},
		QueryString: query,
	}

	resp, err := athenaCli.cli.StartQueryExecution(input)
	if err != nil {
		return nil, err
	}

	return resp.QueryExecutionId, nil
}

func (athenaCli *AthenaClient) GetQueryExecution(queryExecutionID *string) (*athena.QueryExecutionStatus, error) {
	input := &athena.GetQueryExecutionInput{
		QueryExecutionId: queryExecutionID,
	}

	resp, err := athenaCli.cli.GetQueryExecution(input)
	if err != nil {
		return nil, err
	}

	return resp.QueryExecution.Status, nil
}

func (athenaCli *AthenaClient) GetQueryResults(queryExecutionID, nextToken *string) (*athena.ResultSet, *string, error) {
	input := &athena.GetQueryResultsInput{
		QueryExecutionId: queryExecutionID,
	}
	if nextToken != nil {
		input.NextToken = nextToken
	}

	resp, err := athenaCli.cli.GetQueryResults(input)
	if err != nil {
		return nil, nil, err
	}

	return resp.ResultSet, resp.NextToken, nil
}

func (athenaCli *AthenaClient) HandleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case athena.ErrCodeInternalServerException:
			fmt.Println(athena.ErrCodeInternalServerException, aerr.Error())
		case athena.ErrCodeInvalidRequestException:
			fmt.Println(athena.ErrCodeInvalidRequestException, aerr.Error())
		case athena.ErrCodeTooManyRequestsException:
			fmt.Println(athena.ErrCodeTooManyRequestsException, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}
