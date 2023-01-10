package nrsuite

import (
	"bufio"
	"github.com/node-real/nr-test-core/src/log"
	"github.com/node-real/nr-test-core/src/utils"
	"github.com/stretchr/testify/suite"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
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
	IsSkip     bool
}

var util = utils.Utils{}

func parseTestTagInfos(suiteName string) []TagInfo {
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
		log.Error(err)
		return nil
	}
	var tagInfos []TagInfo
	var currFilePath string
	for _, d := range direction {
		//log.Info("package", k)
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
									Line:   fset.File(f.Pos()).Line(cLine.Pos()),
								}
								tagInfos = append(tagInfos, tagInfo)
							}
						}
					}
				}
			}
		}
	}
	return parseTagInfo(suiteName, tagInfos, currFilePath)
}

func parseTagInfo(targetSuite string, tagInfos []TagInfo, filePath string) []TagInfo {
	file, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	var newTagInfos []TagInfo
	for fileScanner.Scan() {
		for _, tagInfo := range tagInfos {
			if lineCount == tagInfo.Line+1 {
				testNameLine := fileScanner.Text()
				if strings.HasPrefix(testNameLine, "type ") {
					suiteName := util.GetStringInBetween(testNameLine, "type ", " struct")
					if suiteName == targetSuite {
						tagInfo.SuiteName = suiteName
						tagInfo.IsSuite = true
						newTagInfos = append(newTagInfos, tagInfo)
					}
				} else if strings.HasPrefix(testNameLine, "func") {
					compileRegex := regexp.MustCompile("\\)(.*?)\\(")
					matchArr := compileRegex.FindStringSubmatch(testNameLine)
					var methodName string
					if matchArr != nil && len(matchArr) > 0 {
						methodName = matchArr[len(matchArr)-1]
					} else {
						break
					}

					suiteRegex := regexp.MustCompile("\\*(.*?)\\)")
					matchArr1 := suiteRegex.FindStringSubmatch(testNameLine)
					var suiteName string
					if matchArr1 != nil && len(matchArr1) > 0 {
						suiteName = matchArr1[len(matchArr1)-1]
					} else {
						break
					}

					methodName = strings.Trim(methodName, " ")
					if methodName != "" && strings.HasPrefix(methodName, "Test") && suiteName == targetSuite {
						tagInfo.MethodName = methodName
						tagInfo.SuiteName = suiteName
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

func parseTagToCaseInfo(tag TagInfo) *suite.CaseInfo {
	tagMap := tag.TagMap
	if tag.IsSuite {
		return nil
	}
	caseInfo := suite.CaseInfo{}
	caseInfo.MethodName = tag.MethodName
	caseInfo.SuiteName = tag.SuiteName
	caseInfo.TagStr = tag.TagStr
	caseInfo.IsSkip = tag.IsSkip
	if tagMap != nil {
		caseInfo.DataKey = tagMap["$RunDataKey"]
		parallelCount := tagMap["$ParallelCount"]
		//retryCount := tagMap["$RetryCount"]
		if parallelCount != "" {
			count, err := util.ConvertStrToInt(parallelCount)
			if err != nil {
				caseInfo.ParallelCount = 1
			} else {
				caseInfo.ParallelCount = count
			}
		}
		//if retryCount != "" {
		//	count, err := util.ConvertStrToInt(retryCount)
		//	if err != nil {
		//		caseInfo.ParallelCount = 1
		//	} else {
		//		caseInf
		//	}
		//}
	}
	return &caseInfo
}
