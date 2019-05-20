package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

type CloudTrailClient struct {
	cli *cloudtrail.CloudTrail
}

func NewCloudTrail(sess *session.Session) *CloudTrailClient {
	client := cloudtrail.New(sess)

	return &CloudTrailClient{cli: client}
}

func (ct *CloudTrailClient) DescribeTrails(input *cloudtrail.DescribeTrailsInput) *cloudtrail.DescribeTrailsOutput {
	resp, err := ct.cli.DescribeTrails(input)
	if err != nil {
		fmt.Println("Got error describing trails:")
		fmt.Println(err.Error())
	}

	return resp
}
