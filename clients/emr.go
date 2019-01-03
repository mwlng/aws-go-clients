package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/emr"
)

type EMRClient struct {
    cli *emr.EMR
}

func NewEMR(sess *session.Session) *EMRClient {
    client := emr.New(sess)

    return &EMRClient{ cli: client }
}

func (emrCli *EMRClient) ListClusters(states []*string) *[]*emr.ClusterSummary {
    input := &emr.ListClustersInput{
       ClusterStates: states,
    }
    resp, err := emrCli.cli.ListClusters(input)
    if err != nil {
        handleError(err)
    }
    clusters := resp.Clusters
    for resp.Marker != nil {
        input = &emr.ListClustersInput{ Marker: resp.Marker }
        resp, err = emrCli.cli.ListClusters(input)
        if err != nil {
            emrCli.handleError(err)
        }
        clusters = append(clusters, resp.Clusters...)
    }
    return &clusters
}

func (emrCli *EMRClient) DescribeCluster(id *string) *emr.DescribeClusterOutput {
    input := &emr.DescribeClusterInput{
        ClusterId: id,
    }
    resp, err := emrCli.cli.DescribeCluster(input)
    if err != nil {
        handleError(err)
    }
    return resp
}

func (emrCli *EMRClient) handleError(err error) {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case emr.ErrCodeInternalServerException:
            fmt.Println(emr.ErrCodeInternalServerException, aerr.Error())
        case emr.ErrCodeInvalidRequestException:
            fmt.Println(emr.ErrCodeInvalidRequestException, aerr.Error())
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        fmt.Println(err.Error())
    }
}
