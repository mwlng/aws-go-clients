package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBClient struct {
	cli *dynamodb.DynamoDB
}

func NewDynamoDB(sess *session.Session) *DynamoDBClient {
	client := dynamodb.New(sess)

	return &DynamoDBClient{cli: client}
}

func (dynamoDBCli *DynamoDBClient) CreateTable(tableName *string,
	attributeDefinitions []*dynamodb.AttributeDefinition,
	keySchema []*dynamodb.KeySchemaElement,
	provisionedThroughput *dynamodb.ProvisionedThroughput,
) error {
	input := &dynamodb.CreateTableInput{
		TableName:             tableName,
		AttributeDefinitions:  attributeDefinitions,
		KeySchema:             keySchema,
		ProvisionedThroughput: provisionedThroughput,
	}

	_, err := dynamoDBCli.cli.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

func (dynamoDBCli *DynamoDBClient) ListTables() ([]*string, error) {
	input := &dynamodb.ListTablesInput{}
	result, err := dynamoDBCli.cli.ListTables(input)
	if err != nil {
		return nil, err
	}
	return result.TableNames, nil
}

func (dynamoDBCli *DynamoDBClient) GetItem(tableName *string,
	key map[string]*dynamodb.AttributeValue, item interface{}) error {
	input := &dynamodb.GetItemInput{
		TableName: tableName,
		Key:       key,
	}
	result, err := dynamoDBCli.cli.GetItem(input)
	if err != nil {
		return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil
	}
	return nil
}

func (dynamoDBCli *DynamoDBClient) PutItem(tableName *string,
	key map[string]*dynamodb.AttributeValue, item interface{}) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: tableName,
		Item:      av,
	}
	_, err = dynamoDBCli.cli.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (dynamoDBCli *DynamoDBClient) UpdateItem(tableName *string,
	key map[string]*dynamodb.AttributeValue,
	attributeValues map[string]*dynamodb.AttributeValue) error {
	input := &dynamodb.UpdateItemInput{
		TableName:                 tableName,
		Key:                       key,
		ExpressionAttributeValues: attributeValues,
		ReturnValues:              aws.String("UPDATED_NEW"),
		UpdateExpression:          aws.String("set Rating = :r"),
	}
	_, err := dynamoDBCli.cli.UpdateItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (dynamoDBCli *DynamoDBClient) DeleteItem(tableName *string,
	key map[string]*dynamodb.AttributeValue) error {
	input := &dynamodb.DeleteItemInput{
		TableName: tableName,
		Key:       key,
	}
	_, err := dynamoDBCli.cli.DeleteItem(input)
	if err != nil {
		return err
	}
	return nil
}

func HandleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case dynamodb.ErrCodeInternalServerError:
			fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
	return
}
