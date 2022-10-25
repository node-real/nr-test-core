package testdata

import (
	"bufio"
	"github.com/node-real/nr-test-core/src/log"
	"os"
)

type DataOperator struct {
}

func (dataOp *DataOperator) GetOnlineData(key string) string {
	//TODO: Robert
	return ""
}

func (dataOp *DataOperator) UploadOnlineData(key string) bool {
	//TODO: Robert
	return false
}

func (dataOp *DataOperator) ReadCsvData() []string {
	//TODO: Robert
	return nil
}

func (dataOp *DataOperator) ReadCustomFileData() map[string]string {
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

func (dataOp *DataOperator) GenerateGql() string {
	//TODO: Robert
	return ""
}
