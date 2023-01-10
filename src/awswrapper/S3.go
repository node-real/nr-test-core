package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/node-real/nr-test-core/src/log"
	"os"
)

var bucketName = "tf-nodereal-qa-test-center"

func getSession() *session.Session {
	cred := getAwsCredentials()
	akId := cred.AccessKeyId
	ak := cred.SecretAccessKey
	sToken := cred.SessionToken
	awsEnv := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsEnv),
		Credentials: credentials.NewStaticCredentials(*akId, *ak, *sToken),
	}))
	return sess
}

func UploadFileToS3(fileName string, s3Key string) {
	file, err := os.Open(fileName)
	if err == nil {
		log.Error(err)
	}
	uploader := s3manager.NewUploader(getSession())
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(s3Key),
		Body:   file,
	})

	if err != nil {
		log.Errorf("Failed upload file '%s' to s3.", fileName, err)
	} else {
		log.Info("Succeed to upload file to s3.")
	}
}
