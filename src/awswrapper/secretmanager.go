package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/node-real/nr-test-core/src/log"
)

func GetSecretValue(secretKey string) string {
	if secretKey == "" {
		log.Warnf("Invoke GetSecretValue method, but the param secretKey is empty")
		return ""
	}
	cred := getAwsCredentials()
	akId := cred.AccessKeyId
	ak := cred.SecretAccessKey
	sToken := cred.SessionToken
	awsEnv := GetAwsRegion()

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsEnv),
		Credentials: credentials.NewStaticCredentials(*akId, *ak, *sToken),
	}))

	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretKey),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				log.Error(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				log.Error(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				log.Error(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeDecryptionFailure:
				log.Error(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				log.Error(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				log.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Error(err.Error())
		}
		return ""
	}

	return *result.SecretString
}
