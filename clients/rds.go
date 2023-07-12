package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
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

func (rdsCli *RDSClient) CreateClusterSnapshot(clusterID, snapshotID string, tags []*rds.Tag) *rds.DBClusterSnapshot {
	input := &rds.CreateDBClusterSnapshotInput{
		DBClusterIdentifier:         aws.String(clusterID),
		DBClusterSnapshotIdentifier: aws.String(snapshotID),
		Tags:                        tags,
	}

	resp, err := rdsCli.cli.CreateDBClusterSnapshot(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	return resp.DBClusterSnapshot
}

func (rdsCli *RDSClient) CopyClusterSnapshot(region, srcSnapshotID, tgtSnapshotID, kmsKeyID string) *rds.DBClusterSnapshot {
	var input *rds.CopyDBClusterSnapshotInput
	if kmsKeyID == "" {
		input = &rds.CopyDBClusterSnapshotInput{
			CopyTags:                          aws.Bool(true),
			DestinationRegion:                 aws.String(region),
			SourceDBClusterSnapshotIdentifier: aws.String(srcSnapshotID),
			TargetDBClusterSnapshotIdentifier: aws.String(tgtSnapshotID),
		}
	} else {
		input = &rds.CopyDBClusterSnapshotInput{
			CopyTags:                          aws.Bool(true),
			DestinationRegion:                 aws.String(region),
			SourceDBClusterSnapshotIdentifier: aws.String(srcSnapshotID),
			TargetDBClusterSnapshotIdentifier: aws.String(tgtSnapshotID),
			KmsKeyId:                          aws.String(kmsKeyID),
		}
	}

	resp, err := rdsCli.cli.CopyDBClusterSnapshot(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	return resp.DBClusterSnapshot
}

func (rdsCli *RDSClient) DescribeClusterSnapshot(clusterID, snapshotID string) *rds.DBClusterSnapshot {
	input := &rds.DescribeDBClusterSnapshotsInput{
		DBClusterIdentifier:         aws.String(clusterID),
		DBClusterSnapshotIdentifier: aws.String(snapshotID),
	}

	resp, err := rdsCli.cli.DescribeDBClusterSnapshots(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	if len(resp.DBClusterSnapshots) > 0 {
		return resp.DBClusterSnapshots[0]
	}

	return nil
}

func (rdsCli *RDSClient) DeleteClusterSnapshot(snapshotID string) *rds.DeleteDBClusterSnapshotOutput {
	input := &rds.DeleteDBClusterSnapshotInput{
		DBClusterSnapshotIdentifier: aws.String(snapshotID),
	}

	resp, err := rdsCli.cli.DeleteDBClusterSnapshot(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	return resp
}

func (rdsCli *RDSClient) CreateDBInstance(input *rds.CreateDBInstanceInput) *rds.DBInstance {
	resp, err := rdsCli.cli.CreateDBInstance(input)
	if err != nil {
		rdsCli.handleError(err)
		return nil
	}

	return resp.DBInstance
}

func (rdsCli *RDSClient) DescribeClusterDBInstances(dbClusterID string) []*rds.DBInstance {
	input := &rds.DescribeDBInstancesInput{
		Filters: append([]*rds.Filter{},
			&rds.Filter{
				Name:   aws.String("db-cluster-id"),
				Values: []*string{aws.String(dbClusterID)},
			},
		),
	}

	resp, err := rdsCli.cli.DescribeDBInstances(input)
	if err != nil {
		rdsCli.handleError(err)
		return nil
	}

	return resp.DBInstances
}

func (rdsCli *RDSClient) DescribeDBInstance(dbInstanceID string) *rds.DBInstance {
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(dbInstanceID),
	}

	resp, err := rdsCli.cli.DescribeDBInstances(input)
	if err != nil {
		rdsCli.handleError(err)
		return nil
	}

	return resp.DBInstances[0]
}

func (rdsCli *RDSClient) CreateDBSnapshot(instanceID, snapshotID string, tags []*rds.Tag) *rds.DBSnapshot {
	input := &rds.CreateDBSnapshotInput{
		DBInstanceIdentifier: aws.String(instanceID),
		DBSnapshotIdentifier: aws.String(snapshotID),
		Tags:                 tags,
	}

	resp, err := rdsCli.cli.CreateDBSnapshot(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	return resp.DBSnapshot
}

func (rdsCli *RDSClient) CopyDBSnapshot(region, srcSnapshotID, tgtSnapshotID, kmsKeyID string) *rds.DBSnapshot {
	var input *rds.CopyDBSnapshotInput
	if kmsKeyID == "" {
		input = &rds.CopyDBSnapshotInput{
			CopyTags:                   aws.Bool(true),
			DestinationRegion:          aws.String(region),
			SourceDBSnapshotIdentifier: aws.String(srcSnapshotID),
			TargetDBSnapshotIdentifier: aws.String(tgtSnapshotID),
		}
	} else {
		input = &rds.CopyDBSnapshotInput{
			CopyTags:                   aws.Bool(true),
			DestinationRegion:          aws.String(region),
			SourceDBSnapshotIdentifier: aws.String(srcSnapshotID),
			TargetDBSnapshotIdentifier: aws.String(tgtSnapshotID),
			KmsKeyId:                   aws.String(kmsKeyID),
		}
	}

	resp, err := rdsCli.cli.CopyDBSnapshot(input)

	if err != nil {
		rdsCli.handleError(err)
	}

	return resp.DBSnapshot
}

func (rdsCli *RDSClient) DescribeDBSnapshot(instanceID, snapshotID string) *rds.DBSnapshot {
	input := &rds.DescribeDBSnapshotsInput{
		DBInstanceIdentifier: aws.String(instanceID),
		DBSnapshotIdentifier: aws.String(snapshotID),
	}

	resp, err := rdsCli.cli.DescribeDBSnapshots(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	if len(resp.DBSnapshots) > 0 {
		return resp.DBSnapshots[0]
	}

	return nil
}

func (rdsCli *RDSClient) DeleteDBSnapshot(snapshotID string) *rds.DeleteDBSnapshotOutput {
	input := &rds.DeleteDBSnapshotInput{
		DBSnapshotIdentifier: aws.String(snapshotID),
	}

	resp, err := rdsCli.cli.DeleteDBSnapshot(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	return resp
}

func (rdsCli *RDSClient) DescribeDBCluster(dbClusterIdentifier string) *rds.DBCluster {
	input := &rds.DescribeDBClustersInput{
		DBClusterIdentifier: &dbClusterIdentifier,
	}

	resp, err := rdsCli.cli.DescribeDBClusters(input)
	if err != nil {
		rdsCli.handleError(err)
		return nil
	}

	return resp.DBClusters[0]
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

func (rdsCli *RDSClient) ListAllDBClusterSnapshots(snapshotType string) []*rds.DBClusterSnapshot {
	input := &rds.DescribeDBClusterSnapshotsInput{
		SnapshotType: aws.String(snapshotType),
	}

	resp, err := rdsCli.cli.DescribeDBClusterSnapshots(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	snapshots := resp.DBClusterSnapshots

	for resp.Marker != nil {
		input = &rds.DescribeDBClusterSnapshotsInput{
			SnapshotType: aws.String(snapshotType),
			Marker:       resp.Marker,
		}

		resp, err := rdsCli.cli.DescribeDBClusterSnapshots(input)
		if err != nil {
			rdsCli.handleError(err)
		}

		snapshots = append(snapshots, resp.DBClusterSnapshots...)
	}

	return snapshots
}

func (rdsCli *RDSClient) ListDBClusterSnapshots(clusterID, snapshotType string) []*rds.DBClusterSnapshot {
	input := &rds.DescribeDBClusterSnapshotsInput{
		DBClusterIdentifier: aws.String(clusterID),
		SnapshotType:        aws.String(snapshotType),
	}

	resp, err := rdsCli.cli.DescribeDBClusterSnapshots(input)
	if err != nil {
		rdsCli.handleError(err)
	}

	snapshots := resp.DBClusterSnapshots

	for resp.Marker != nil {
		input = &rds.DescribeDBClusterSnapshotsInput{
			DBClusterIdentifier: aws.String(clusterID),
			SnapshotType:        aws.String(snapshotType),
			Marker:              resp.Marker,
		}

		resp, err := rdsCli.cli.DescribeDBClusterSnapshots(input)
		if err != nil {
			rdsCli.handleError(err)
		}

		snapshots = append(snapshots, resp.DBClusterSnapshots...)
	}

	return snapshots
}

func (rdsCli *RDSClient) DeleteCluster(clusterID, finalSnapshotID string) *rds.DeleteDBClusterOutput {
	input := &rds.DeleteDBClusterInput{
		DBClusterIdentifier: aws.String(clusterID),
		SkipFinalSnapshot:   aws.Bool(true),
	}
	if len(finalSnapshotID) > 0 {
		input = &rds.DeleteDBClusterInput{
			DBClusterIdentifier:       aws.String(clusterID),
			FinalDBSnapshotIdentifier: aws.String(finalSnapshotID),
			SkipFinalSnapshot:         aws.Bool(false),
		}
	}

	result, err := rdsCli.cli.DeleteDBCluster(input)
	if err != nil {
		handleError(err)
	}

	return result
}

func (rdsCli *RDSClient) RestoreDClusterFromSnapshot(input *rds.RestoreDBClusterFromSnapshotInput) *rds.DBCluster {
	resp, err := rdsCli.cli.RestoreDBClusterFromSnapshot(input)
	if err != nil {
		rdsCli.handleError(err)
		return nil
	}

	return resp.DBCluster
}

func (rdsCli *RDSClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case rds.ErrCodeDBSnapshotAlreadyExistsFault:
			fmt.Println(rds.ErrCodeDBSnapshotAlreadyExistsFault, aerr.Error())
		case rds.ErrCodeDBSnapshotNotFoundFault:
			fmt.Println(rds.ErrCodeDBSnapshotNotFoundFault, aerr.Error())
		case rds.ErrCodeInvalidDBSnapshotStateFault:
			fmt.Println(rds.ErrCodeInvalidDBSnapshotStateFault, aerr.Error())
		case rds.ErrCodeSnapshotQuotaExceededFault:
			fmt.Println(rds.ErrCodeSnapshotQuotaExceededFault, aerr.Error())
		case rds.ErrCodeKMSKeyNotAccessibleFault:
			fmt.Println(rds.ErrCodeKMSKeyNotAccessibleFault, aerr.Error())
		case rds.ErrCodeCustomAvailabilityZoneNotFoundFault:
			fmt.Println(rds.ErrCodeCustomAvailabilityZoneNotFoundFault, aerr.Error())
		case rds.ErrCodeDBClusterAlreadyExistsFault:
			fmt.Println(rds.ErrCodeDBClusterAlreadyExistsFault, aerr.Error())
		case rds.ErrCodeDBClusterQuotaExceededFault:
			fmt.Println(rds.ErrCodeDBClusterQuotaExceededFault, aerr.Error())
		case rds.ErrCodeStorageQuotaExceededFault:
			fmt.Println(rds.ErrCodeStorageQuotaExceededFault, aerr.Error())
		case rds.ErrCodeDBSubnetGroupNotFoundFault:
			fmt.Println(rds.ErrCodeDBSubnetGroupNotFoundFault, aerr.Error())
		case rds.ErrCodeDBClusterSnapshotNotFoundFault:
			fmt.Println(rds.ErrCodeDBClusterSnapshotNotFoundFault, aerr.Error())
		case rds.ErrCodeInsufficientDBClusterCapacityFault:
			fmt.Println(rds.ErrCodeInsufficientDBClusterCapacityFault, aerr.Error())
		case rds.ErrCodeInsufficientStorageClusterCapacityFault:
			fmt.Println(rds.ErrCodeInsufficientStorageClusterCapacityFault, aerr.Error())
		case rds.ErrCodeInvalidDBClusterSnapshotStateFault:
			fmt.Println(rds.ErrCodeInvalidDBClusterSnapshotStateFault, aerr.Error())
		case rds.ErrCodeInvalidVPCNetworkStateFault:
			fmt.Println(rds.ErrCodeInvalidVPCNetworkStateFault, aerr.Error())
		case rds.ErrCodeInvalidRestoreFault:
			fmt.Println(rds.ErrCodeInvalidRestoreFault, aerr.Error())
		case rds.ErrCodeInvalidSubnet:
			fmt.Println(rds.ErrCodeInvalidSubnet, aerr.Error())
		case rds.ErrCodeOptionGroupNotFoundFault:
			fmt.Println(rds.ErrCodeOptionGroupNotFoundFault, aerr.Error())
		case rds.ErrCodeDomainNotFoundFault:
			fmt.Println(rds.ErrCodeDomainNotFoundFault, aerr.Error())
		case rds.ErrCodeDBClusterParameterGroupNotFoundFault:
			fmt.Println(rds.ErrCodeDBClusterParameterGroupNotFoundFault, aerr.Error())
		case rds.ErrCodeInvalidDBInstanceStateFault:
			fmt.Println(rds.ErrCodeInvalidDBInstanceStateFault, aerr.Error())
		case rds.ErrCodeInvalidDBClusterStateFault:
			fmt.Println(rds.ErrCodeInvalidDBClusterStateFault, aerr.Error())
		case rds.ErrCodeInvalidDBSubnetGroupStateFault:
			fmt.Println(rds.ErrCodeInvalidDBSubnetGroupStateFault, aerr.Error())
		case rds.ErrCodeDBClusterNotFoundFault:
			fmt.Println(rds.ErrCodeDBClusterNotFoundFault, aerr.Error())
		case rds.ErrCodeDBInstanceNotFoundFault:
			fmt.Println(rds.ErrCodeDBInstanceNotFoundFault, aerr.Error())
		case rds.ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs:
			fmt.Println(rds.ErrCodeDBSubnetGroupDoesNotCoverEnoughAZs, aerr.Error())
		case rds.ErrCodeGlobalClusterNotFoundFault:
			fmt.Println(rds.ErrCodeGlobalClusterNotFoundFault, aerr.Error())
		case rds.ErrCodeInvalidGlobalClusterStateFault:
			fmt.Println(rds.ErrCodeInvalidGlobalClusterStateFault, aerr.Error())
		case rds.ErrCodeDBInstanceAlreadyExistsFault:
			fmt.Println(rds.ErrCodeDBInstanceAlreadyExistsFault, aerr.Error())
		case rds.ErrCodeInsufficientDBInstanceCapacityFault:
			fmt.Println(rds.ErrCodeInsufficientDBInstanceCapacityFault, aerr.Error())
		case rds.ErrCodeDBParameterGroupNotFoundFault:
			fmt.Println(rds.ErrCodeDBParameterGroupNotFoundFault, aerr.Error())
		case rds.ErrCodeDBSecurityGroupNotFoundFault:
			fmt.Println(rds.ErrCodeDBSecurityGroupNotFoundFault, aerr.Error())
		case rds.ErrCodeInstanceQuotaExceededFault:
			fmt.Println(rds.ErrCodeInstanceQuotaExceededFault, aerr.Error())
		case rds.ErrCodeProvisionedIopsNotAvailableInAZFault:
			fmt.Println(rds.ErrCodeProvisionedIopsNotAvailableInAZFault, aerr.Error())
		case rds.ErrCodeStorageTypeNotSupportedFault:
			fmt.Println(rds.ErrCodeStorageTypeNotSupportedFault, aerr.Error())
		case rds.ErrCodeAuthorizationNotFoundFault:
			fmt.Println(rds.ErrCodeAuthorizationNotFoundFault, aerr.Error())
		case rds.ErrCodeBackupPolicyNotFoundFault:
			fmt.Println(rds.ErrCodeBackupPolicyNotFoundFault, aerr.Error())
		case rds.ErrCodeCertificateNotFoundFault:
			fmt.Println(rds.ErrCodeCertificateNotFoundFault, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}
