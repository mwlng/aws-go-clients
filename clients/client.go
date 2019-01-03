package clients

import (
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/iam"
    "github.com/aws/aws-sdk-go/service/emr"
    "github.com/aws/aws-sdk-go/service/ecr"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/cloudformation"
)

func NewClient(service string, sess *session.Session) interface{} {
    var client interface{}
    switch service {
        case "ec2":
            client = &EC2Client{ cli: ec2.New(sess) }
        case "iam":
            client = &IAMClient{ cli: iam.New(sess) }
        case "s3":
            client = &S3Client{ cli: s3.New(sess) }
        case "emr":
            client = &EMRClient{ cli: emr.New(sess) }
        case "ecr":
            client = &ECRClient{ cli: ecr.New(sess) }
        case "cloudformation":
            client = &CfnClient{ cli: cloudformation.New(sess) }
        default:
            client = nil
    }
    return client
}
