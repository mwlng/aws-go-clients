package clients

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/sts"
)

type STSClient struct {
    cli *sts.STS
}

func NewSTS(sess *session.Session) *STSClient {
    client := sts.New(sess)

    return &STSClient{ cli: client }
}

func (stsCli *STSClient) GetSessionCredsWithMfa(mfaSN *string, tokenCode *string, duration *int64) *sts.Credentials {
    input := &sts.GetSessionTokenInput {
        SerialNumber: mfaSN,
        TokenCode: tokenCode,
        DurationSeconds: duration,
    }
    resp, err := stsCli.cli.GetSessionToken(input)
    if err != nil {
        stsCli.handleError(err)
        return nil
    }
    return resp.Credentials
}

func (stsCli *STSClient) handleError(err error) {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case sts.ErrCodeRegionDisabledException:
            fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
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
