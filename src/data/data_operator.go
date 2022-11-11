package data

import (
	"bufio"
	"encoding/csv"
	"github.com/node-real/nr-test-core/src/awswrapper"
	"github.com/node-real/nr-test-core/src/log"
	"os"
	"strings"
)

type DataOperator struct {
}

func (dataOp *DataOperator) GetSecretData(key string) string {
	return awswrapper.GetSecretValue(key)
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

//func (dataOp *DataOperator) ReadCustomFileDatas(dirPath string) []map[string]string {
//	//TODO:Robert
//	return nil
//}

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
