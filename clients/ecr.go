package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type ECRClient struct {
	cli *ecr.ECR
}

func NewECR(sess *session.Session) *ECRClient {
	client := ecr.New(sess)

	return &ECRClient{cli: client}
}

func (ecrCli *ECRClient) CreateRepository(repoName string) {
	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(repoName),
	}

	_, err := ecrCli.cli.CreateRepository(input)

	if err != nil {
		ecrCli.handleError(err)
	}
}

func (ecrCli *ECRClient) ListRepositories() []*ecr.Repository {
	input := &ecr.DescribeRepositoriesInput{}

	resp, err := ecrCli.cli.DescribeRepositories(input)
	if err != nil {
		ecrCli.handleError(err)
	}

	repositories := resp.Repositories
	input = &ecr.DescribeRepositoriesInput{NextToken: resp.NextToken}

	for resp.NextToken != nil {
		resp, err = ecrCli.cli.DescribeRepositories(input)
		if err != nil {
			ecrCli.handleError(err)
		}

		repositories = append(repositories, resp.Repositories...)
	}

	return repositories
}

func (ecrCli *ECRClient) ListImageIdsByRepository(repoName *string) []*ecr.ImageIdentifier {
	input := &ecr.ListImagesInput{RepositoryName: repoName}

	resp, err := ecrCli.cli.ListImages(input)
	if err != nil {
		ecrCli.handleError(err)
	}

	images := resp.ImageIds

	for resp.NextToken != nil {
		input = &ecr.ListImagesInput{
			NextToken:      resp.NextToken,
			RepositoryName: repoName,
		}

		resp, err = ecrCli.cli.ListImages(input)
		if err != nil {
			ecrCli.handleError(err)
		}

		images = append(images, resp.ImageIds...)
	}

	return images
}

func (ecrCli *ECRClient) DescribeImageByID(repoName *string, id *ecr.ImageIdentifier) *ecr.ImageDetail {
	input := &ecr.DescribeImagesInput{
		RepositoryName: repoName,
		ImageIds:       []*ecr.ImageIdentifier{id},
	}

	resp, err := ecrCli.cli.DescribeImages(input)
	if err != nil {
		ecrCli.handleError(err)

		return &ecr.ImageDetail{}
	}

	return resp.ImageDetails[0]
}

func (ecrCli *ECRClient) SetRepositoryPolicy(input *ecr.SetRepositoryPolicyInput) *ecr.SetRepositoryPolicyOutput {
	resp, err := ecrCli.cli.SetRepositoryPolicy(input)
	if err != nil {
		ecrCli.handleError(err)
	}

	return resp
}

func (ecrCli *ECRClient) GetRepositoryPolicy(input *ecr.GetRepositoryPolicyInput) *ecr.GetRepositoryPolicyOutput {
	resp, err := ecrCli.cli.GetRepositoryPolicy(input)
	if err != nil {
		ecrCli.handleError(err)
	}

	return resp
}

func (ecrCli *ECRClient) DeleteRepository(input *ecr.DeleteRepositoryInput) *ecr.DeleteRepositoryOutput {
	resp, err := ecrCli.cli.DeleteRepository(input)
	if err != nil {
		ecrCli.handleError(err)
	}

	return resp
}

func (ecrCli *ECRClient) GetAuthorizationToken() []*ecr.AuthorizationData {
	input := &ecr.GetAuthorizationTokenInput{}

	result, err := ecrCli.cli.GetAuthorizationToken(input)
	if err != nil {
		ecrCli.handleError(err)

		return nil
	}

	return result.AuthorizationData
}

func (ecrCli *ECRClient) UploadImage(srcImage, imageTag, registryID, RepoName string) {
	input := &ecr.PutImageInput{
		ImageManifest:  aws.String(srcImage),
		ImageTag:       aws.String(imageTag),
		RegistryId:     aws.String(registryID),
		RepositoryName: aws.String(RepoName),
	}

	_, err := ecrCli.cli.PutImage(input)
	if err != nil {
		ecrCli.handleError(err)
	}
}

func (ecrCli *ECRClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case ecr.ErrCodeServerException:
			fmt.Println(ecr.ErrCodeServerException, aerr.Error())
		case ecr.ErrCodeInvalidParameterException:
			fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
		case ecr.ErrCodeInvalidTagParameterException:
			fmt.Println(ecr.ErrCodeInvalidTagParameterException, aerr.Error())
		case ecr.ErrCodeTooManyTagsException:
			fmt.Println(ecr.ErrCodeTooManyTagsException, aerr.Error())
		case ecr.ErrCodeRepositoryAlreadyExistsException:
			fmt.Println(ecr.ErrCodeRepositoryAlreadyExistsException, aerr.Error())
		case ecr.ErrCodeLimitExceededException:
			fmt.Println(ecr.ErrCodeLimitExceededException, aerr.Error())
		case ecr.ErrCodeKmsException:
			fmt.Println(ecr.ErrCodeKmsException, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}
