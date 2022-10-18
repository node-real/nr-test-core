package nrsuite

import (
	"bufio"
	"fmt"
	"github.com/node-real/nr-test-core/src/utils"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type TagInfo struct {
	TagStr     string
	TagMap     map[string]string
	Line       int
	SuiteName  string
	MethodName string
	IsSuite    bool
}

func parseTestTagInfos() []TagInfo {
	var filePaths []string
	for m := 0; m < 5; m++ {
		_, filename, _, r := runtime.Caller(m)
		if !r || strings.HasSuffix(filename, "/testing/testing.go") {
			break
		} else {
			filePaths = append(filePaths, filename)
		}
	}
	fset := token.NewFileSet() // positions are relative to fset
	direction, err := parser.ParseDir(fset, "./", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var tagInfos []TagInfo
	var currFilePath string
	for k, d := range direction {
		fmt.Println("package", k)
		for n, f := range d.Files {
			for _, fileP := range filePaths {
				if filepath.Base(fileP) == n {
					currFilePath = fileP
					for _, c := range f.Comments {
						for _, cLine := range c.List {
							if strings.HasPrefix(cLine.Text, "// Tags::") || strings.HasPrefix(cLine.Text, "//Tags::") {
								tagInfo := TagInfo{
									TagStr: cLine.Text,
									TagMap: parseTagStr(cLine.Text),
									Line:   fset.File(f.Pos()).Line(c.Pos()),
								}
								tagInfos = append(tagInfos, tagInfo)
							}
						}
					}
				}
			}
		}
	}
	return parseTagInfo(tagInfos, currFilePath)
}

func parseTagInfo(tagInfos []TagInfo, filePath string) []TagInfo {
	file, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	var newTagInfos []TagInfo
	var currSuite string
	for fileScanner.Scan() {
		for _, tagInfo := range tagInfos {
			if lineCount == tagInfo.Line+1 {
				testNameLine := fileScanner.Text()
				if strings.HasPrefix(testNameLine, "type ") {
					suiteName := utils.GetStringInBetween(testNameLine, "type ", " struct")
					tagInfo.SuiteName = suiteName
					tagInfo.IsSuite = true
					currSuite = suiteName
					newTagInfos = append(newTagInfos, tagInfo)
				} else if strings.HasPrefix(testNameLine, "func") {
					methodName := utils.GetStringInBetween(testNameLine, ")", "()")
					methodName = strings.Trim(methodName, " ")
					if strings.HasPrefix(methodName, "Test") {
						tagInfo.MethodName = methodName
						tagInfo.SuiteName = currSuite
						tagInfo.IsSuite = false
						newTagInfos = append(newTagInfos, tagInfo)
					}
				}
			}
		}
		lineCount++
	}
	defer file.Close()
	return newTagInfos
}

func parseTagStr(tagStr string) map[string]string {
	tagStr = strings.TrimLeft(tagStr, "// Tags:: ")
	tags := strings.Split(tagStr, " ")
	tagMap := map[string]string{}
	for _, tagItem := range tags {
		tagIn := strings.Split(tagItem, ":")
		if len(tagIn) == 2 {
			tagMap[tagIn[0]] = tagIn[1]
		}
	}
	return tagMap
}
