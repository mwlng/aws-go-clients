package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudformation"
    "github.com/aws/aws-sdk-go/aws/awserr"
)

type CFNClient struct {
    cli *cloudformation.CloudFormation
}

func NewCloudformation(sess *session.Session) *CFNClient {
    client := cloudformation.New(sess)

    return &CFNClient{ cli: client }
}

func (cfn *CFNClient) ListStacks() *[]*cloudformation.StackSummary {
    input := &cloudformation.ListStacksInput{}
    resp, err := cfn.cli.ListStacks(input)
    if err != nil {
       cfn.handleError(err)
    }
    summaries := resp.StackSummaries
    for resp.NextToken != nil {
        input = &cloudformation.ListStacksInput{ NextToken: resp.NextToken }
        resp, err = cfn.cli.ListStacks(input)
        if err != nil {
            cfn.handleError(err)
        }
        summaries = append(summaries, resp.StackSummaries...)
    }
    return &summaries
}

func (cfn *CFNClient) GetTemplate(stackName *string) *string {
    input := &cloudformation.GetTemplateInput{ 
        StackName: stackName,
    }
    resp, err := cfn.cli.GetTemplate(input)
    if err != nil {
       cfn.handleError(err)
    }
    return resp.TemplateBody
}

func (cfn *CFNClient) ListStackResources(stackName *string) *[]*cloudformation.StackResourceSummary {
    input := &cloudformation.ListStackResourcesInput{
        StackName: stackName,
    }
    resp, err := cfn.cli.ListStackResources(input)
    if err != nil {
        cfn.handleError(err)
    }
    summaries := resp.StackResourceSummaries
    for resp.NextToken != nil {
        input = &cloudformation.ListStackResourcesInput{ NextToken: resp.NextToken }
        resp, err = cfn.cli.ListStackResources(input)
        if err != nil {
            cfn.handleError(err)
        }
        summaries = append(summaries, resp.StackResourceSummaries...)
    }
    return &summaries
}

func (cfn *CFNClient) ListChangeSets(stackName *string) *[]*cloudformation.ChangeSetSummary {
    input := &cloudformation.ListChangeSetsInput{
        StackName: stackName,
    }
    resp, err := cfn.cli.ListChangeSets(input)
    if err != nil {
        cfn.handleError(err)
    }
    summaries := resp.Summaries
    for resp.NextToken != nil {
        input = &cloudformation.ListChangeSetsInput{ NextToken: resp.NextToken }
        resp, err = cfn.cli.ListChangeSets(input)
        if err != nil {
            cfn.handleError(err)
        }
        summaries = append(summaries, resp.Summaries...)
    }
    return &summaries
}

func (cfn *CFNClient) ListStackSets() *[]*cloudformation.StackSetSummary  {
    input := &cloudformation.ListStackSetsInput{}
    resp, err := cfn.cli.ListStackSets(input)
    if err != nil {
        cfn.handleError(err)
    }
    summaries := resp.Summaries
    for resp.NextToken != nil {
        input = &cloudformation.ListStackSetsInput{ NextToken: resp.NextToken }
        resp, err = cfn.cli.ListStackSets(input)
        if err != nil {
            cfn.handleError(err)
        }
        summaries = append(summaries, resp.Summaries...)
    }
    return &summaries
}

func (cfn *CFNClient) handleError(err error) {
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
