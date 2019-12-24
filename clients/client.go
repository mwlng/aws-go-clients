package clients

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewClient(service string, sess *session.Session) interface{} {
	var client interface{}
	switch service {
	case "athena":
		client = NewAthena(sess)
	case "autoscaling":
		client = NewASG(sess)
	case "cloudformation":
		client = NewCloudformation(sess)
	case "ec2":
		client = NewEC2(sess)
	case "ecs":
		client = NewECS(sess)
	case "ecr":
		client = NewECR(sess)
	case "emr":
		client = NewEMR(sess)
	case "glue":
		client = NewGlue(sess)
	case "iam":
		client = NewIAM(sess)
	case "route53":
		client = NewR53(sess)
	case "redshift":
		client = NewRedShift(sess)
	case "s3":
		client = NewS3(sess)
	case "sts":
		client = NewSTS(sess)
	default:
		client = nil
	}
	return client
}
