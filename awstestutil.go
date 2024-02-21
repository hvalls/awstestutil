package awstestutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func getSession() *session.Session {
	sess, err := session.NewSession(
		aws.NewConfig().WithRegion("eu-west-1").WithCredentials(credentials.AnonymousCredentials).WithEndpoint("http://localhost:4566"))
	if err != nil {
		panic(err)
	}
	return sess
}

func GetIAMRole(name string) *iam.Role {
	sess := getSession()
	svc := iam.New(sess)
	input := &iam.GetRoleInput{
		RoleName: aws.String(name),
	}
	r, err := svc.GetRole(input)
	if err != nil {
		return nil
	}
	return r.Role
}

func GetIAMInstanceProfile(name string) *iam.InstanceProfile {
	sess := getSession()
	svc := iam.New(sess)
	input := &iam.GetInstanceProfileInput{
		InstanceProfileName: aws.String(name),
	}
	r, err := svc.GetInstanceProfile(input)
	if err != nil {
		return nil
	}
	return r.InstanceProfile
}
