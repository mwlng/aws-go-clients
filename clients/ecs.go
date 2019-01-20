package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ecs"
)

type ECSClient struct{
    cli *ecs.ECS
}

func NewECS(sess *session.Session) *ECSClient {
    client := ecs.New(sess)

    return &ECSClient{ cli: client }
}

func (ecsCli *ECSClient) ListAllClusters() *[]*ecs.Cluster {
    input := &ecs.ListClustersInput{} 
    resp, err := ecsCli.cli.ListClusters(input)
    if err != nil {
        ecsCli.handleError(err)
    }
    clusterArns := resp.ClusterArns
    for resp.NextToken != nil {
        input = &ecs.ListClustersInput{ NextToken: resp.NextToken, }
        resp, err := ecsCli.cli.ListClusters(input)
        if err != nil {
            ecsCli.handleError(err)
        }
        clusterArns = append(clusterArns, resp.ClusterArns...)
    }
    if len(clusterArns) <= 0 { return nil }

    return ecsCli.describeClusters(&clusterArns) 
}

func (ecsCli *ECSClient) describeClusters(clusterArns *[]*string) *[]*ecs.Cluster {
    total := len(*clusterArns)
    if total <= 0 { return nil }

    i := 0
    batchStart := i * 100
    batchEnd := (i+1) * 100
    if batchEnd > total {
       batchEnd = total % 100
    }
    input := &ecs.DescribeClustersInput{ Clusters: (*clusterArns)[batchStart:batchEnd], }
    resp, err := ecsCli.cli.DescribeClusters(input)
    if err != nil {
         ecsCli.handleError(err)
    }
    clusters := resp.Clusters
    i = 1
    for {
        if i*100 >= total { break }
        batchEnd = (i+1) * 100
        if batchEnd > total {
            batchEnd = i * 100 + total % 100
        }

        input = &ecs.DescribeClustersInput{ Clusters: (*clusterArns)[batchStart:batchEnd], }
        resp, err = ecsCli.cli.DescribeClusters(input)
        if err != nil {
            ecsCli.handleError(err)
        }
        clusters = append(clusters, resp.Clusters...)
        i += 1
    }
    return &clusters
}

func (ecsCli *ECSClient) ListAllServices() *[]*ecs.Service {
    input := &ecs.ListServicesInput{} 
    resp, err := ecsCli.cli.ListServices(input)
    if err != nil {
        ecsCli.handleError(err)
    }
    serviceArns := resp.ServiceArns
    for resp.NextToken != nil {
        input = &ecs.ListServicesInput{ NextToken: resp.NextToken }
        resp, err = ecsCli.cli.ListServices(input)
        if err != nil {
            ecsCli.handleError(err)
        }
        serviceArns = append(serviceArns, resp.ServiceArns...)
    }
    
    if len(serviceArns) <= 0 { return nil } 
    return nil
}

func (ecsCli *ECSClient) ListAllTasks() *[]*ecs.Task {
    input := &ecs.ListTasksInput{}
    resp, err := ecsCli.cli.ListTasks(input)
    if err != nil {
        ecsCli.handleError(err)
    }
    taskArns := resp.TaskArns
    for resp.NextToken != nil {
        input = &ecs.ListTasksInput{ NextToken: resp.NextToken }
        resp, err = ecsCli.cli.ListTasks(input)
        if err != nil {
            ecsCli.handleError(err)
        }
        taskArns = append(taskArns, resp.TaskArns...)
    }
    if len(taskArns) <=0 { return nil }
    return nil
}

func (ecsCli *ECSClient) ListAllTaskDefinitions() *[]*ecs.TaskDefinition {
    input := &ecs.ListTaskDefinitionsInput{}
    resp, err := ecsCli.cli.ListTaskDefinitions(input)
    if err != nil {
        ecsCli.handleError(err)
    }
    definitionArns := resp.TaskDefinitionArns
    for resp.NextToken != nil {
        input = &ecs.ListTaskDefinitionsInput{ NextToken: resp.NextToken }
        resp, err = ecsCli.cli.ListTaskDefinitions(input)
        if err != nil {
            ecsCli.handleError(err)
        }
        definitionArns = append(definitionArns, resp.TaskDefinitionArns...)
    }

    if len(definitionArns) <=0 { return nil } 

    return nil
}

func (ecsCli *ECSClient) handleError(err error) {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case ecs.ErrCodeServerException:
            fmt.Println(ecs.ErrCodeServerException, aerr.Error())
        case ecs.ErrCodeClientException:
            fmt.Println(ecs.ErrCodeClientException, aerr.Error())
        case ecs.ErrCodeInvalidParameterException:
            fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
        case ecs.ErrCodeClusterNotFoundException:
            fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
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
