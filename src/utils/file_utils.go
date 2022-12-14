package utils

import (
	"bufio"
	"github.com/node-real/nr-test-core/src/log"
	"os"
)

//var log = log.Log

type FileUtils struct {
}

func (utils *FileUtils) ReadFileToLines(filePath string) []string {
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

func (utils *FileUtils) WriteFileWithLines(dataItems []string, filePath string) {
	var file *os.File
	var err1 error
	if utils.CheckFileIsExist(filePath) {
		file, err1 = os.OpenFile(filePath, os.O_APPEND, 0666)
	} else {
		file, err1 = os.Create(filePath)
	}
	for _, item := range dataItems {
		file.WriteString(item)
		file.WriteString("\n")
	}
	if err1 != nil {
		log.Fatal(err1)
	}
	defer file.Close()
	file.Sync()
}

func (utils *FileUtils) CheckFileIsExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
