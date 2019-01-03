package service

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/ecr"
)

func HandleError(err error) {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case ecr.ErrCodeServerException:
            fmt.Println(ecr.ErrCodeServerException, aerr.Error())
        case ecr.ErrCodeInvalidParameterException:
            fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
        case ecr.ErrCodeRepositoryNotFoundException:
            fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
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
