package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Client struct{
    cli *ec2.EC2
}

func NewEC2(sess *session.Session) *EC2Client {
    client := ec2.New(sess)

    return &EC2Client{ cli: client }
}

func (ec2cli *EC2Client) ListAllVpcs() []*ec2.Vpc {
    input := &ec2.DescribeVpcsInput{} 
    resp, err := ec2cli.cli.DescribeVpcs(input)
    if err != nil {
        fmt.Println("Got error:")
        fmt.Println(err.Error())
    }

    return resp.Vpcs
}

func (ec2cli *EC2Client) ListAllAvailbleZones() *ec2.DescribeAvailabilityZonesOutput {
    return nil
}

func (ec2cli *EC2Client) ListAllSubnets() *ec2.DescribeSubnetsOutput  {
    return  nil
}

func (ec2cli *EC2Client) ListAMIsByOwner(owner string) *ec2.DescribeImagesOutput {
    input := &ec2.DescribeImagesInput{
        Owners: []*string{
            aws.String(owner),
        },
    }
    resp, err := ec2cli.cli.DescribeImages(input)
    if err != nil {
        ec2cli.handleError(err)
    }

    return resp
}

func (ec2cli *EC2Client) handleError(err error) {
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
