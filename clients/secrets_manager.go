package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretsManagerClient struct {
	cli *secretsmanager.SecretsManager
}

func NewSecretsManager(sess *session.Session) *SecretsManagerClient {
	client := secretsmanager.New(sess)

	return &SecretsManagerClient{cli: client}
}

func (smCli *SecretsManagerClient) GetSecret(name string) string {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(name),
		VersionId:    nil,
		VersionStage: nil,
	}

	resp, err := smCli.cli.GetSecretValue(input)
	if err != nil {
		smCli.handleError(err)

		return ""
	}

	return *resp.SecretString
}

func (smCli *SecretsManagerClient) CreateSecret(name, value string) {
	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(name),
		SecretString: aws.String(value),
	}

	_, err := smCli.cli.CreateSecret(input)
	if err != nil {
		smCli.handleError(err)
	}
}

func (smCli *SecretsManagerClient) PutSecret(name, value string) {
	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(name),
		SecretString: aws.String(value),
	}

	_, err := smCli.cli.PutSecretValue(input)
	if err != nil {
		smCli.handleError(err)
	}
}

func (smCli *SecretsManagerClient) UpdateSecret(name, value string) {
	input := &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(name),
		SecretString: aws.String(value),
	}

	_, err := smCli.cli.UpdateSecret(input)
	if err != nil {
		smCli.handleError(err)
	}
}

func (smCli *SecretsManagerClient) ListAllSecrets() []*secretsmanager.SecretListEntry {
	secrets := []*secretsmanager.SecretListEntry{}
	input := &secretsmanager.ListSecretsInput{}
	resp, err := smCli.cli.ListSecrets(input)
	if err != nil {
		smCli.handleError(err)
		return secrets
	}

	secrets = append(secrets, resp.SecretList...)

	for resp.NextToken != nil {
		input = &secretsmanager.ListSecretsInput{
			NextToken: resp.NextToken,
		}

		resp, err = smCli.cli.ListSecrets(input)
		if err != nil {
			smCli.handleError(err)
			return secrets
		}

		secrets = append(secrets, resp.SecretList...)
	}

	return secrets
}

func (smCli *SecretsManagerClient) handleError(err error) {

	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case secretsmanager.ErrCodeInvalidParameterException:
			fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
		case secretsmanager.ErrCodeInvalidRequestException:
			fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
		case secretsmanager.ErrCodeLimitExceededException:
			fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
		case secretsmanager.ErrCodeEncryptionFailure:
			fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
		case secretsmanager.ErrCodeResourceExistsException:
			fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
		case secretsmanager.ErrCodeResourceNotFoundException:
			fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
		case secretsmanager.ErrCodeInternalServiceError:
			fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
