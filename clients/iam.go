package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type IAMClient struct {
	cli *iam.IAM
}

func NewIAM(sess *session.Session) *IAMClient {
	client := iam.New(sess)

	return &IAMClient{cli: client}
}

func (iamCli *IAMClient) ListUsers() []*iam.User {
	input := &iam.ListUsersInput{}

	resp, err := iamCli.cli.ListUsers(input)
	if err != nil {
		iamCli.handleError(err)
	}

	users := resp.Users

	for *resp.IsTruncated {
		input = &iam.ListUsersInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListUsers(input)
		if err != nil {
			iamCli.handleError(err)
		}

		users = append(users, resp.Users...)
	}

	return users
}

func (iamCli *IAMClient) GetUserPolicy(userName *string, policyName *string) *string {
	input := &iam.GetUserPolicyInput{
		UserName:   userName,
		PolicyName: policyName,
	}

	resp, err := iamCli.cli.GetUserPolicy(input)
	if err != nil {
		iamCli.handleError(err)
	}

	return resp.PolicyDocument
}

func (iamCli *IAMClient) ListUserPolicies(userName *string) []*string {
	input := &iam.ListUserPoliciesInput{UserName: userName}

	resp, err := iamCli.cli.ListUserPolicies(input)
	if err != nil {
		iamCli.handleError(err)
	}

	policyNames := resp.PolicyNames

	for *resp.IsTruncated {
		input = &iam.ListUserPoliciesInput{
			Marker:   resp.Marker,
			UserName: userName,
		}

		resp, err = iamCli.cli.ListUserPolicies(input)
		if err != nil {
			iamCli.handleError(err)
		}

		policyNames = append(policyNames, resp.PolicyNames...)
	}

	return policyNames
}

func (iamCli *IAMClient) ListAttachedUserPolicies(userName *string) []*iam.AttachedPolicy {
	input := &iam.ListAttachedUserPoliciesInput{
		UserName: userName,
	}

	resp, err := iamCli.cli.ListAttachedUserPolicies(input)
	if err != nil {
		iamCli.handleError(err)
	}

	attachedPolicies := resp.AttachedPolicies

	for *resp.IsTruncated {
		input = &iam.ListAttachedUserPoliciesInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListAttachedUserPolicies(input)
		if err != nil {
			iamCli.handleError(err)
		}

		attachedPolicies = append(attachedPolicies, resp.AttachedPolicies...)
	}

	return attachedPolicies
}

func (iamCli *IAMClient) ListGroupsForUser(userName *string) []*iam.Group {
	input := &iam.ListGroupsForUserInput{
		UserName: userName,
	}

	resp, err := iamCli.cli.ListGroupsForUser(input)
	if err != nil {
		iamCli.handleError(err)
	}

	return resp.Groups
}

func (iamCli *IAMClient) ListGroups() []*iam.Group {
	input := &iam.ListGroupsInput{}

	resp, err := iamCli.cli.ListGroups(input)
	if err != nil {
		iamCli.handleError(err)
	}

	groups := resp.Groups

	for *resp.IsTruncated {
		input = &iam.ListGroupsInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListGroups(input)
		if err != nil {
			iamCli.handleError(err)
		}

		groups = append(groups, resp.Groups...)
	}

	return groups
}

func (iamCli *IAMClient) ListGroupPolicies(groupName *string) []*string {
	input := &iam.ListGroupPoliciesInput{GroupName: groupName}

	resp, err := iamCli.cli.ListGroupPolicies(input)
	if err != nil {
		iamCli.handleError(err)
	}

	policyNames := resp.PolicyNames

	for *resp.IsTruncated {
		input = &iam.ListGroupPoliciesInput{
			Marker:    resp.Marker,
			GroupName: groupName,
		}

		resp, err = iamCli.cli.ListGroupPolicies(input)
		if err != nil {
			iamCli.handleError(err)
		}

		policyNames = append(policyNames, resp.PolicyNames...)
	}

	return policyNames
}

func (iamCli *IAMClient) GetGroupPolicy(groupName *string, policyName *string) *string {
	input := &iam.GetGroupPolicyInput{
		GroupName:  groupName,
		PolicyName: policyName,
	}

	resp, err := iamCli.cli.GetGroupPolicy(input)
	if err != nil {
		iamCli.handleError(err)
	}

	return resp.PolicyDocument
}

func (iamCli *IAMClient) ListAttachedGroupPolicies(groupName *string) []*iam.AttachedPolicy {
	input := &iam.ListAttachedGroupPoliciesInput{
		GroupName: groupName,
	}

	resp, err := iamCli.cli.ListAttachedGroupPolicies(input)
	if err != nil {
		iamCli.handleError(err)
	}

	attachedPolicies := resp.AttachedPolicies

	for *resp.IsTruncated {
		input = &iam.ListAttachedGroupPoliciesInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListAttachedGroupPolicies(input)
		if err != nil {
			iamCli.handleError(err)
		}

		attachedPolicies = append(attachedPolicies, resp.AttachedPolicies...)
	}

	return attachedPolicies
}

func (iamCli *IAMClient) ListRoles() []*iam.Role {
	input := &iam.ListRolesInput{}

	resp, err := iamCli.cli.ListRoles(input)
	if err != nil {
		iamCli.handleError(err)
	}

	roles := resp.Roles

	for *resp.IsTruncated {
		input = &iam.ListRolesInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListRoles(input)
		if err != nil {
			iamCli.handleError(err)
		}

		roles = append(roles, resp.Roles...)
	}

	return roles
}

func (iamCli *IAMClient) ListRolePolicies(roleName *string) []*string {
	input := &iam.ListRolePoliciesInput{
		RoleName: roleName,
	}

	resp, err := iamCli.cli.ListRolePolicies(input)
	if err != nil {
		iamCli.handleError(err)
	}

	policyNames := resp.PolicyNames

	for *resp.IsTruncated {
		input = &iam.ListRolePoliciesInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListRolePolicies(input)
		if err != nil {
			iamCli.handleError(err)
		}

		policyNames = append(policyNames, resp.PolicyNames...)
	}

	return policyNames
}

func (iamCli *IAMClient) GetRolePolicy(roleName *string, policyName *string) *string {
	input := &iam.GetRolePolicyInput{
		RoleName:   roleName,
		PolicyName: policyName,
	}

	resp, err := iamCli.cli.GetRolePolicy(input)
	if err != nil {
		iamCli.handleError(err)
	}

	return resp.PolicyDocument
}

func (iamCli *IAMClient) ListAttachedRolePolicies(roleName *string) []*iam.AttachedPolicy {
	input := &iam.ListAttachedRolePoliciesInput{
		RoleName: roleName,
	}

	resp, err := iamCli.cli.ListAttachedRolePolicies(input)
	if err != nil {
		iamCli.handleError(err)
	}

	attachedPolicies := resp.AttachedPolicies

	for *resp.IsTruncated {
		input = &iam.ListAttachedRolePoliciesInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListAttachedRolePolicies(input)
		if err != nil {
			iamCli.handleError(err)
		}

		attachedPolicies = append(attachedPolicies, resp.AttachedPolicies...)
	}

	return attachedPolicies
}

func (iamCli *IAMClient) GetRole(name *string) *iam.Role {
	input := &iam.GetRoleInput{
		RoleName: name,
	}

	resp, err := iamCli.cli.GetRole(input)
	if err != nil {
		iamCli.handleError(err)

		return nil
	}

	return resp.Role
}

func (iamCli *IAMClient) ListPolicies() []*iam.Policy {
	input := &iam.ListPoliciesInput{}

	resp, err := iamCli.cli.ListPolicies(input)
	if err != nil {
		iamCli.handleError(err)
	}

	policies := resp.Policies

	for *resp.IsTruncated {
		input = &iam.ListPoliciesInput{Marker: resp.Marker}

		resp, err = iamCli.cli.ListPolicies(input)
		if err != nil {
			iamCli.handleError(err)
		}

		policies = append(policies, resp.Policies...)
	}

	return policies
}

func (iamCli *IAMClient) GetPolicyVersion(policyArn *string, verID *string) *iam.PolicyVersion {
	input := &iam.GetPolicyVersionInput{
		PolicyArn: policyArn,
		VersionId: verID,
	}

	resp, err := iamCli.cli.GetPolicyVersion(input)
	if err != nil {
		iamCli.handleError(err)
	}

	return resp.PolicyVersion
}

func (iamCli *IAMClient) GetPolicy(policyArn *string) *iam.GetPolicyOutput {
	input := &iam.GetPolicyInput{
		PolicyArn: policyArn,
	}

	resp, err := iamCli.cli.GetPolicy(input)
	if err != nil {
		iamCli.handleError(err)
	}

	return resp
}

func (iamCli *IAMClient) CreateRole(name *string, path *string, assumeRolePolicyDocument *string) (*iam.Role, error) {
	input := &iam.CreateRoleInput{
		RoleName:                 name,
		Path:                     path,
		AssumeRolePolicyDocument: assumeRolePolicyDocument,
	}

	resp, err := iamCli.cli.CreateRole(input)
	if err != nil {
		iamCli.handleError(err)

		return nil, err
	}

	return resp.Role, nil
}

func (iamCli *IAMClient) DeleteRole(name *string) error {
	input := &iam.DeleteRoleInput{
		RoleName: name,
	}

	_, err := iamCli.cli.DeleteRole(input)
	if err != nil {
		iamCli.handleError(err)

		return err
	}

	return nil
}

func (iamCli *IAMClient) AttachRolePolicy(roleName *string, policyArn *string) error {
	input := &iam.AttachRolePolicyInput{
		RoleName:  roleName,
		PolicyArn: policyArn,
	}

	_, err := iamCli.cli.AttachRolePolicy(input)
	if err != nil {
		iamCli.handleError(err)

		return err
	}

	return nil
}

func (iamCli *IAMClient) DetachRolePolicy(roleName *string, policyArn *string) error {
	input := &iam.DetachRolePolicyInput{
		RoleName:  roleName,
		PolicyArn: policyArn,
	}

	_, err := iamCli.cli.DetachRolePolicy(input)
	if err != nil {
		iamCli.handleError(err)

		return err
	}

	return nil
}

func (iamCli *IAMClient) handleError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case iam.ErrCodeLimitExceededException:
			fmt.Println(iam.ErrCodeLimitExceededException, aerr.Error())
		case iam.ErrCodeInvalidInputException:
			fmt.Println(iam.ErrCodeInvalidInputException, aerr.Error())
		case iam.ErrCodeEntityAlreadyExistsException:
			fmt.Println(iam.ErrCodeEntityAlreadyExistsException, aerr.Error())
		case iam.ErrCodeDeleteConflictException:
			fmt.Println(iam.ErrCodeDeleteConflictException, aerr.Error())
		case iam.ErrCodeNoSuchEntityException:
			fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
		case iam.ErrCodeMalformedPolicyDocumentException:
			fmt.Println(iam.ErrCodeMalformedPolicyDocumentException, aerr.Error())
		case iam.ErrCodeUnmodifiableEntityException:
			fmt.Println(iam.ErrCodeUnmodifiableEntityException, aerr.Error())
		case iam.ErrCodeConcurrentModificationException:
			fmt.Println(iam.ErrCodeConcurrentModificationException, aerr.Error())
		case iam.ErrCodePolicyNotAttachableException:
			fmt.Println(iam.ErrCodePolicyNotAttachableException, aerr.Error())
		case iam.ErrCodeServiceFailureException:
			fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
