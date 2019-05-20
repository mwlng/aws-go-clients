package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Client struct {
	cli *ec2.EC2
}

func NewEC2(sess *session.Session) *EC2Client {
	client := ec2.New(sess)

	return &EC2Client{cli: client}
}

func (ec2Cli *EC2Client) ListAllVpcs() []*ec2.Vpc {
	input := &ec2.DescribeVpcsInput{}
	resp, err := ec2Cli.cli.DescribeVpcs(input)
	if err != nil {
		ec2Cli.handleError(err)
	}

	return resp.Vpcs
}

func (ec2Cli *EC2Client) ListAllAvailbleZones() *ec2.DescribeAvailabilityZonesOutput {
	return nil
}

func (ec2Cli *EC2Client) ListAllSubnets() *ec2.DescribeSubnetsOutput {
	return nil
}

func (ec2Cli *EC2Client) ListAllInstances() []*ec2.Instance {
	input := &ec2.DescribeInstancesInput{}
	resp, err := ec2Cli.cli.DescribeInstances(input)
	if err != nil {
		ec2Cli.handleError(err)
	}
	reservations := resp.Reservations
	for resp.NextToken != nil {
		input = &ec2.DescribeInstancesInput{NextToken: resp.NextToken}
		resp, err = ec2Cli.cli.DescribeInstances(input)
		if err != nil {
			ec2Cli.handleError(err)
		}
		reservations = append(reservations, resp.Reservations...)
	}

	var instances []*ec2.Instance
	for _, r := range reservations {
		instances = append(instances, r.Instances...)
	}
	return instances
}

func (ec2Cli *EC2Client) ListAMIsByOwner(owner string) *ec2.DescribeImagesOutput {
	input := &ec2.DescribeImagesInput{
		Owners: []*string{
			aws.String(owner),
		},
	}
	resp, err := ec2Cli.cli.DescribeImages(input)
	if err != nil {
		ec2Cli.handleError(err)
	}

	return resp
}

func (ec2Cli *EC2Client) handleError(err error) {
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
