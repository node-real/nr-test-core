package nrsuite

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/checker"
	"github.com/node-real/nr-test-core/src/core"
	"github.com/node-real/nr-test-core/src/core/nrdriver"
	"github.com/node-real/nr-test-core/src/log"
	"github.com/stretchr/testify/suite"
	"reflect"
	"strings"
	"sync"
	"testing"
)

type NRBaseSuite struct {
	suite.Suite
	nrdriver.TestDriver
	TestData   map[string][]string
	ResultData map[string][]string
	TestName   string
	mu         sync.Mutex
}

func (baseSuite *NRBaseSuite) SetupSuite() {
	log.Infof("=== Setup Test Suite: %s ===", baseSuite.TestName)
	baseSuite.TestDriver = *nrdriver.Driver()
}

func (baseSuite *NRBaseSuite) TearDownSuite() {
	log.Infof("=== TearDown Test Suite: %s ===", baseSuite.TestName)
}

func (baseSuite *NRBaseSuite) BeforeTest(suiteName, testName string) {
	log.Infof("=== Before Test: %s ===", testName)
}

func (baseSuite *NRBaseSuite) AfterTest(suiteName, testName string) {
	log.Infof("=== After Test: %s ===", testName)
}

func (baseSuite *NRBaseSuite) AppendResultData(key string, valueItem string) {
	baseSuite.mu.Lock()
	if baseSuite.ResultData == nil {
		baseSuite.ResultData = map[string][]string{}
	}
	baseSuite.ResultData[key] = append(baseSuite.ResultData[key], valueItem)
	baseSuite.mu.Unlock()
}

func Run(t *testing.T, testSuite suite.TestingSuite) {
	log.Infof("------------Run Test: %s-------------", t.Name())
	core.InitConfig()

	argMap := core.Config.TestFilters

	suiteType := reflect.TypeOf(testSuite)
	currSuiteName := suiteType.Elem().Name()

	tagInfos := parseTestTagInfos(currSuiteName)

	var currSuiteInfo TagInfo
	var skipCases []string
	isSkipSuite := false
	hasFilter := len(argMap) > 0
	hasSuiteTag := false
	currSuiteValue := reflect.ValueOf(testSuite).Elem()
	currSuiteValue.FieldByName("TestName").Set(reflect.ValueOf(t.Name()))
	methodNum := reflect.TypeOf(testSuite).NumMethod()

	for index, tagInfo := range tagInfos {
		if currSuiteName != tagInfo.SuiteName {
			break
		}
		if tagInfo.IsSuite {
			currSuiteInfo = tagInfo
			hasSuiteTag = true
			//Suite Skip Check
			if tagInfo.TagMap["skip"] == "true" {
				isSkipSuite = true
				break
			} else {
				for k, v := range argMap {
					targetTagStr := tagInfo.TagMap[k]
					if targetTagStr == "" {
						tagInfo.IsSkip = true
						isSkipSuite = true
						break
					} else {
						tagValues := strings.Split(targetTagStr, ",")
						argValues := strings.Split(v, ",")
						checker := new(checker.Checker)
						isContainOneV := false
						for _, aValue := range argValues {
							if checker.IsContainsInArray(tagValues, aValue) {
								isContainOneV = true
								break
							}
						}
						if !isContainOneV {
							tagInfo.IsSkip = true
							isSkipSuite = true
						}
					}
				}
			}
		} else {
			//Method Skip Check
			if tagInfo.TagMap["skip"] == "true" {
				skipCases = append(skipCases, tagInfo.MethodName)
				tagInfo.IsSkip = true
			} else {
				for k, v := range argMap {
					targetTagStr := tagInfo.TagMap[k]
					if targetTagStr == "" {
						// if curr tag is null, will use the suite tag, so break
						break
					} else {
						tagValues := strings.Split(targetTagStr, ",")
						argValues := strings.Split(v, ",")
						checker := new(checker.Checker)
						isContainOneV := false
						for _, aValue := range argValues {
							if checker.IsContainsInArray(tagValues, aValue) {
								isContainOneV = true
								break
							}
						}
						if !isContainOneV {
							skipCases = append(skipCases, tagInfo.MethodName)
							tagInfo.IsSkip = true
						}
					}
				}
			}
		}
		tagInfos[index] = tagInfo
	}

	if isSkipSuite {
		fmt.Println("Current suite tag string:", currSuiteInfo.TagStr)
		fmt.Println("Current suite is skipped!")
		log.Info("Current suite tag string:", currSuiteInfo.TagStr)
		log.Infof("Current suite is skipped!")
		t.Skipf("Current Suite Tags:%s", currSuiteInfo.TagStr) // skip current test suite
	}

	// for the no tag test methods, add empty tag info
	for i := 1; i < methodNum; i++ {
		methodName := suiteType.Method(i).Name
		hasMethod := false
		for _, tagI := range tagInfos {
			if tagI.MethodName == methodName {
				hasMethod = true
				break
			}
		}
		if !hasMethod && strings.HasPrefix(methodName, "Test") {
			isSkipM := false
			if hasFilter && !hasSuiteTag {
				isSkipM = true
				skipCases = append(skipCases, methodName)
			}
			tagInfos = append(tagInfos, TagInfo{
				SuiteName:  currSuiteName,
				MethodName: methodName,
				IsSuite:    false,
				IsSkip:     isSkipM, //if test method and suite no tags and has filter will skip
			})
		}
	}
	fmt.Println("Skipped cases list:", skipCases)
	log.Info("Skipped cases list:", skipCases)

	var caseInfos []suite.CaseInfo
	for _, tag := range tagInfos {
		caseInfo := parseTagToCaseInfo(tag)
		if caseInfo != nil {
			caseInfos = append(caseInfos, *caseInfo)
		}
	}
	log.Info("caseInfos: ", caseInfos)
	suite.Run(t, testSuite, caseInfos)
}
