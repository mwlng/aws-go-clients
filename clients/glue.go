package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/glue"
)

type GlueClient struct {
    cli *glue.Glue
}

func NewGlue(sess *session.Session) *GlueClient {
    client := glue.New(sess)

    return &GlueClient{ cli: client }
}

func (glueCli *GlueClient) ListDatabases() *[]*glue.Database {
    input := &glue.GetDatabasesInput{}
    resp, err := glueCli.cli.GetDatabases(input)
    if err != nil {
        glueCli.handleError(err)
    }
    databases := resp.DatabaseList
    for resp.NextToken != nil {
        input = &glue.GetDatabasesInput{ NextToken: resp.NextToken }
        resp, err = glueCli.cli.GetDatabases(input)
        if err != nil {
            glueCli.handleError(err)
        }
        databases = append(databases, resp.DatabaseList...)
    }
    return &databases
}

func (glueCli *GlueClient) ListTables(dbName *string) *[]*glue.Table {
    input := &glue.GetTablesInput{ 
         DatabaseName: dbName,
    }
    resp, err := glueCli.cli.GetTables(input)
    if err != nil {
        glueCli.handleError(err)
    }
    tables := resp.TableList
    for resp.NextToken != nil {
        input = &glue.GetTablesInput{ 
            DatabaseName: dbName,
            NextToken: resp.NextToken,
        }
        resp, err = glueCli.cli.GetTables(input)
        if err != nil {
            glueCli.handleError(err)
        }
        tables = append(tables, resp.TableList...)
    }
    return &tables
}

func (glueCli *GlueClient) ListCrawlers() *[]*glue.Crawler {
    input := &glue.GetCrawlersInput{}
    resp, err := glueCli.cli.GetCrawlers(input)
    if err != nil {
        glueCli.handleError(err)
    }
    crawlers := resp.Crawlers
    
    for resp.NextToken != nil {
        input = &glue.GetCrawlersInput{ NextToken: resp.NextToken }
        resp, err = glueCli.cli.GetCrawlers(input)
        if err != nil {
            glueCli.handleError(err)
        }
        crawlers = append(crawlers, resp.Crawlers...)
    }
    return &crawlers
}

func (glueCli *GlueClient) ListClassifiers() *[]*glue.Classifier {
    input := &glue.GetClassifiersInput{}
    resp, err := glueCli.cli.GetClassifiers(input)
    if err != nil {
        glueCli.handleError(err)
    }
    classifiers := resp.Classifiers
    for resp.NextToken != nil {
        input = &glue.GetClassifiersInput{ NextToken: resp.NextToken }
        resp, err = glueCli.cli.GetClassifiers(input)
        if err != nil {
            glueCli.handleError(err)
        }
        classifiers = append(classifiers, resp.Classifiers...)
    }
    return &classifiers
}

func (glueCli *GlueClient) ListTriggers() *[]*glue.Trigger {
    input := &glue.GetTriggersInput{}
    resp, err := glueCli.cli.GetTriggers(input)
    if err != nil {
        glueCli.handleError(err)
    }
    triggers := resp.Triggers
    for resp.NextToken != nil {
        input = &glue.GetTriggersInput{ NextToken: resp.NextToken }
        resp, err = glueCli.cli.GetTriggers(input)
        if err != nil {
            glueCli.handleError(err)
        }
        triggers = append(triggers, resp.Triggers...)
    }
    return &triggers
}

func (glueCli *GlueClient) handleError(err error) {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case glue.ErrCodeInvalidInputException:
            fmt.Println(glue.ErrCodeInvalidInputException, aerr.Error())
        case glue.ErrCodeEntityNotFoundException:
            fmt.Println(glue.ErrCodeEntityNotFoundException, aerr.Error())
        case glue.ErrCodeInternalServiceException:
            fmt.Println(glue.ErrCodeInternalServiceException, aerr.Error())
        case glue.ErrCodeOperationTimeoutException:
            fmt.Println(glue.ErrCodeOperationTimeoutException, aerr.Error())
        case glue.ErrCodeEncryptionException:
            fmt.Println(glue.ErrCodeEncryptionException,  aerr.Error())
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
