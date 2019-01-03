package service

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
)

type Service struct {
    Profile string
    Region  string
    AccessKey string
    SecretKey string
    SessToken string
    Session *session.Session 
}

func (svc *Service) NewSession() *session.Session {

    var awsConfig *aws.Config 
    var sessOptions session.Options
    var creds *credentials.Credentials
    var value credentials.Value

    if svc.Region != "" {
        awsConfig = &aws.Config{Region: aws.String(svc.Region)}
    } 

    if svc.AccessKey != "" && svc.SecretKey !="" {
        if svc.SessToken != "" {
            value = credentials.Value {
                AccessKeyID:  svc.AccessKey,
                SecretAccessKey: svc.SecretKey,
                SessionToken: svc.SessToken, 
 
            }
        } else {
            value = credentials.Value {
                AccessKeyID:  svc.AccessKey,
                SecretAccessKey: svc.SecretKey,
            }
        }
        creds = credentials.NewStaticCredentialsFromCreds(value)
        if awsConfig != nil {
            awsConfig = awsConfig.WithCredentials(creds)
        } else {
            awsConfig = aws.NewConfig().WithCredentials(creds)
        }
        sessOptions = session.Options{
            Config: *awsConfig,
        }
    } else if svc.Profile != "" {
       if awsConfig != nil {
           sessOptions = session.Options{
               Config: *awsConfig,
               Profile: svc.Profile,
           }
       } else {
           sessOptions = session.Options{
               Profile: svc.Profile,
           }
       }
    } else {
        sessOptions = session.Options{
            SharedConfigState: session.SharedConfigEnable,
        }
    }

    svc.Session = session.Must(session.NewSessionWithOptions(sessOptions))

    return svc.Session
}