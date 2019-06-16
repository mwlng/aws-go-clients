## aws-go-clients

### Overview

aws-go-clients is a golang client library for AWS services. The purpose of this library is to simplify making API call to the AWS services, and hence reduce the number of lines of code in your program.

### Examples

1. Create a client for AWS EC2 service by using  a specific client.
```
import (
	"fmt"

	"github.com/mwlng/aws-go-clients/clients"
	"github.com/mwlng/aws-go-clients/service"
)

func main() {
	svc := service.Service{
		Region:    "us-east-1",
		AccessKey: "<AWS access key id>",
		SecretKey: "<AWS access secret key>",
	}
	sess := svc.NewSession()
	ec2Cli := clients.NewEC2(sess)

	ec2Instances := ec2Cli.ListAllInstances()

	for _, inst := range ec2Instances {
		fmt.Printf("%s = %s\n", *inst.InstanceId, *inst.PrivateIpAddress)
	}
}

```
2. More generic way to create a client for AWS S3 service.
```
package main

import (
	"fmt"

	"github.com/mwlng/aws-go-clients/clients"
	"github.com/mwlng/aws-go-clients/service"
)

func main() {
	svc := service.Service{
		Region:  "us-east-1",
		Profile: "default",
	}
	sess := svc.NewSession()

	s3Cli := clients.NewClient("s3", sess).(*clients.S3Client)

	bucketName := "<s3 bucket name>"
	_, objects := s3Cli.ListObjects(&bucketName, nil, nil)
	for _, obj := range objects {
		fmt.Printf("%s\n", *obj.Key)
	}
}
```