package awswrapper

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"os"
)

func GetAwsRegion() string {
	return os.Getenv("AWS_REGION")
}

func getAwsCredentials() types.Credentials {
	return getCredentials("arn:aws:iam::346509735976:role/tf-nodereal-qa-infra-test-platform-read-role")
}

func getCredentials(roleARN string) types.Credentials {
	ctx := context.TODO()
	awsEnv := "us-east-1"

	var cfg aws.Config
	var err error

	cfg, err = config.LoadDefaultConfig(ctx,
		config.WithRegion(awsEnv),
	)
	if err != nil {
		fmt.Println(err)
	}

	client := sts.NewFromConfig(cfg)
	output, err := client.AssumeRole(ctx, &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleARN),
		RoleSessionName: aws.String("AWSASSUMEROLEARN"),
	})
	if err != nil {
		fmt.Println(err)
	}
	creds := *output.Credentials
	return creds
}
