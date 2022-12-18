package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

type RDSClient struct {
	cli *rds.RDS
}

func NewRDS(sess *session.Session) *RDSClient {
	client := rds.New(sess)

	return &RDSClient{cli: client}
}

func (rdsCli *RDSClient) ListDBInstances() []*rds.DBInstance {
	input := &rds.DescribeDBInstancesInput{}

	resp, err := rdsCli.cli.DescribeDBInstances(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	instances := resp.DBInstances

	for resp.Marker != nil {
		input = &rds.DescribeDBInstancesInput{Marker: resp.Marker}

		resp, err = rdsCli.cli.DescribeDBInstances(input)
		if err != nil {
			rdsCli.handleError(err)
		}

		instances = append(instances, resp.DBInstances...)
	}

	return instances
}

func (rdsCli *RDSClient) ListDBClusters() []*rds.DBCluster {
	input := &rds.DescribeDBClustersInput{}

	resp, err := rdsCli.cli.DescribeDBClusters(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	clusters := resp.DBClusters

	for resp.Marker != nil {
		input = &rds.DescribeDBClustersInput{
			Marker: resp.Marker,
		}

		resp, err = rdsCli.cli.DescribeDBClusters(input)
		if err != nil {
			rdsCli.handleError(err)
		}

		clusters = append(clusters, resp.DBClusters...)
	}

	return clusters
}

func (rdsCli *RDSClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case rds.ErrCodeDBClusterNotFoundFault:
			fmt.Println(rds.ErrCodeDBClusterNotFoundFault, aerr.Error())
		case rds.ErrCodeDBInstanceNotFoundFault:
			fmt.Println(rds.ErrCodeDBInstanceNotFoundFault, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}
