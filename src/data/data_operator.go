package data

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/node-real/nr-test-core/src/log"
	"os"
	"strings"
)

type DataOperator struct {
}

func (dataOp *DataOperator) GetSecretData(secretId string) string {
	roleToAssumeArn := "*****"
	sessionName := "*****"
	regain := os.Getenv("AWS_REGION")
	fmt.Println(regain)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	svc0 := sts.New(sess)
	r, err := svc0.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         &roleToAssumeArn,
		RoleSessionName: &sessionName,
	})
	akId := r.Credentials.AccessKeyId
	ak := r.Credentials.SecretAccessKey
	sToken := r.Credentials.SessionToken

	sess.Config.Credentials = credentials.NewStaticCredentials(*akId, *ak, *sToken)

	svc := secretsmanager.New(sess)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretId),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeDecryptionFailure:
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return ""
	}

	return result.String()
}

func (dataOp *DataOperator) GetOnlineData(key string) string {
	//TODO: Robert
	return ""
}

func (dataOp *DataOperator) UploadOnlineData(key string) bool {
	//TODO: Robert
	return false
}

func (dataOp *DataOperator) ReadCsvData(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	return preData
}

func (dataOp *DataOperator) ReadCustomCaseData(filePath string) []CustomCaseData {
	readFile, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		//fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var cases []CustomCaseData
	var currCase CustomCaseData
	currCaseSepa := ""
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		if strings.HasPrefix(currLine, "--") {
			currCaseSepa = strings.TrimPrefix(currLine, "--")
			switch currCaseSepa {
			case CaseDescFlag:
				currCase = CustomCaseData{
					CaseDesc:  "",
					CaseInfos: map[string]string{},
				}
				break
			case CaseEnd:
				cases = append(cases, currCase)
				currCase = CustomCaseData{}
				currCaseSepa = ""
				break
			default:
				currCase.CaseInfos[currCaseSepa] += ""
				break
			}
		} else {
			if currCase.CaseInfos != nil {
				if currCaseSepa == CaseDescFlag {
					currCase.CaseDesc = currLine
				} else {
					currCase.CaseInfos[currCaseSepa] += currLine
				}
			}
		}
	}
	return cases
}

func (dataOp *DataOperator) ReadCustomFileDatas(dirPath string) []map[string]string {
	//TODO:Robert
	return nil
}

func (dataOp *DataOperator) ReadFileLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
