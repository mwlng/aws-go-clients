package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

type R53Client struct {
	cli *route53.Route53
}

func NewR53(sess *session.Session) *R53Client {
	client := route53.New(sess)

	return &R53Client{cli: client}
}

func (r53Cli *R53Client) ListHostedZones() []*route53.HostedZone {
	input := &route53.ListHostedZonesInput{}

	resp, err := r53Cli.cli.ListHostedZones(input)
	if err != nil {
		r53Cli.handleError(err)
	}

	zones := resp.HostedZones

	for *resp.IsTruncated {
		input = &route53.ListHostedZonesInput{Marker: resp.Marker}

		resp, err = r53Cli.cli.ListHostedZones(input)
		if err != nil {
			r53Cli.handleError(err)
		}

		zones = append(zones, resp.HostedZones...)
	}

	return zones
}

func (r53Cli *R53Client) ListResourceRecordSets(hostedZoneID *string) []*route53.ResourceRecordSet {
	input := &route53.ListResourceRecordSetsInput{HostedZoneId: hostedZoneID}

	resp, err := r53Cli.cli.ListResourceRecordSets(input)
	if err != nil {
		r53Cli.handleError(err)
	}

	records := resp.ResourceRecordSets

	for *resp.IsTruncated {
		input = &route53.ListResourceRecordSetsInput{
			HostedZoneId:    hostedZoneID,
			StartRecordName: resp.NextRecordName,
		}

		resp, err = r53Cli.cli.ListResourceRecordSets(input)
		if err != nil {
			r53Cli.handleError(err)
		}

		records = append(records, resp.ResourceRecordSets...)
	}

	return records
}

func (r53Cli *R53Client) ListGeoLocations() []*route53.GeoLocationDetails {
	input := &route53.ListGeoLocationsInput{}

	resp, err := r53Cli.cli.ListGeoLocations(input)
	if err != nil {
		r53Cli.handleError(err)
	}

	return resp.GeoLocationDetailsList
}

func (r53Cli *R53Client) GetResourceRecordSet(name *string, hostedZoneID *string) *route53.ResourceRecordSet {
	recordSets := r53Cli.ListResourceRecordSets(hostedZoneID)
	for _, recordSet := range recordSets {
		recordSetName := recordSet.Name
		if *recordSetName == *name {
			return recordSet
		}
	}

	return nil
}

func (r53Cli *R53Client) ChangeResourceRecordSets(recordSets []*route53.ResourceRecordSet, action *string,
	hostedZoneID *string, changeComment *string) *route53.ChangeResourceRecordSetsOutput {
	changes := []*route53.Change{}

	for _, recordSet := range recordSets {
		change := &route53.Change{
			Action:            action,
			ResourceRecordSet: recordSet,
		}

		changes = append(changes, change)
	}

	if len(changes) > 0 {
		input := &route53.ChangeResourceRecordSetsInput{
			ChangeBatch: &route53.ChangeBatch{
				Changes: changes,
				Comment: changeComment,
			},
			HostedZoneId: hostedZoneID,
		}

		resp, err := r53Cli.cli.ChangeResourceRecordSets(input)
		if err != nil {
			handleError(err)
		}

		return resp
	}

	return nil
}

func (r53Cli *R53Client) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case route53.ErrCodeNoSuchHostedZone:
			fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
		case route53.ErrCodeNoSuchHealthCheck:
			fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
		case route53.ErrCodeInvalidChangeBatch:
			fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
		case route53.ErrCodeInvalidInput:
			fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
		case route53.ErrCodePriorRequestNotComplete:
			fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
}
