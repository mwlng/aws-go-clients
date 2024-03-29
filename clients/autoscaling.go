package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

type ASGClient struct {
	cli *autoscaling.AutoScaling
}

func NewASG(sess *session.Session) *ASGClient {
	client := autoscaling.New(sess)

	return &ASGClient{cli: client}
}

func (asgCli *ASGClient) DescribeAutoScalngInstances(instanceID string) *autoscaling.DescribeAutoScalingInstancesOutput {
	input := &autoscaling.DescribeAutoScalingInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	result, err := asgCli.cli.DescribeAutoScalingInstances(input)

	if err != nil {
		asgCli.handleError(err)
	}

	return result
}

func (asgCli *ASGClient) GetAutoScalingGroupByName(name string) *autoscaling.Group {
	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{&name},
	}

	resp, err := asgCli.cli.DescribeAutoScalingGroups(input)
	if err != nil {
		asgCli.handleError(err)
	}

	groups := resp.AutoScalingGroups

	if len(groups) > 0 {
		return groups[0]
	}

	return nil
}

func (asgCli *ASGClient) ListAllAutoScalingGroups() []*autoscaling.Group {
	input := &autoscaling.DescribeAutoScalingGroupsInput{}

	resp, err := asgCli.cli.DescribeAutoScalingGroups(input)
	if err != nil {
		asgCli.handleError(err)
	}

	groups := resp.AutoScalingGroups

	for resp.NextToken != nil {
		input = &autoscaling.DescribeAutoScalingGroupsInput{NextToken: resp.NextToken}

		resp, err = asgCli.cli.DescribeAutoScalingGroups(input)
		if err != nil {
			asgCli.handleError(err)
		}

		groups = append(groups, resp.AutoScalingGroups...)
	}

	return groups
}

func (asgCli *ASGClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case autoscaling.ErrCodeInvalidNextToken:
			fmt.Println(autoscaling.ErrCodeInvalidNextToken, aerr.Error())
		case autoscaling.ErrCodeResourceContentionFault:
			fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}
