package nrsuite

import (
	"flag"
	"fmt"
	"github.com/node-real/nr-test-core/src/checker"
	"github.com/node-real/nr-test-core/src/core"
	"github.com/node-real/nr-test-core/src/core/nrdriver"
	"github.com/stretchr/testify/suite"
	"reflect"
	"strings"
	"sync"
	"testing"
)

var (
	once   sync.Once
	config *core.RunningConfig
)

type NRBaseSuite struct {
	suite.Suite
	nrdriver.TestDriver
	TestData   map[string][]string
	ResultData map[string][]string
	mu         sync.Mutex
}

func (baseSuite *NRBaseSuite) SetupSuite() {
	initTest(baseSuite)
}

func (baseSuite *NRBaseSuite) TearDownTestSuite() {
}

func (baseSuite *NRBaseSuite) BeforeTest() {
}

func (baseSuite *NRBaseSuite) AfterTest(suiteName, testName string) {
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
	//nrBaseSuite := testSuite.(NRBaseSuite)
	if !flag.Parsed() {
		flag.Parsed()
	}

	once.Do(func() {
		configV := parseRunningConfig()
		config = &configV
	})

	argMap := config.TestFilters

	tagInfos := parseTestTagInfos()
	currSuiteName := reflect.TypeOf(testSuite).Elem().Name()
	//var skipCases []string
	//var currSuiteTags string
	var currSuiteInfo TagInfo
	isSkipSuite := false
	for _, tagInfo := range tagInfos {
		if currSuiteName != tagInfo.SuiteName {
			break
		}
		if tagInfo.IsSuite {
			currSuiteInfo = tagInfo
			//Suite Skip Check
			if tagInfo.TagMap["skip"] == "true" {
				isSkipSuite = true
				break
			} else {
				for k, v := range argMap {
					targetTagStr := tagInfo.TagMap[k]
					if targetTagStr == "" {
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
							isSkipSuite = true
						}
					}
				}
			}
		} else {
			//Method Skip Check
			if tagInfo.TagMap["skip"] == "true" {
				//skipCases = append(skipCases, tagInfo.MethodName)
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
							//skipCases = append(skipCases, tagInfo.MethodName)
							tagInfo.IsSkip = false
						}
					}
				}
			}
		}
	}

	if isSkipSuite {
		t.Skipf("Current Suite Tags:%s", currSuiteInfo.TagStr) // skip current test suite
	}
	var caseInfos []suite.CaseInfo
	for _, tag := range tagInfos {
		caseInfo := parseTagToCaseInfo(tag)
		if caseInfo != nil {
			caseInfos = append(caseInfos, *caseInfo)
		}
	}
	fmt.Println("caseInfos: ", caseInfos)
	suite.Run(t, testSuite, caseInfos)
}

func initTest(baseSuite *NRBaseSuite) {
	baseSuite.TestDriver = nrdriver.Driver()
	baseSuite.RunningConfig = config
}

//func getSkipCases(tagInfos []TagInfo) []string {
//	var skipCases []string
//	for _, tagInfo := range tagInfos {
//		if tagInfo.TagMap["skip"] == "true" && tagInfo.MethodName != "" {
//			skipCases = append(skipCases, tagInfo.MethodName)
//		}
//	}
//	return skipCases
//}

//reflect.ValueOf(testSuite).FieldByName("TagInfos").(tagInfos)
//suiteValue := reflect.TypeOf(testSuite).Elem()
//baseSuite, r := suiteValue.FieldByName("NRBaseSuite")
//var skipCases []string
//if r {
//	runningTag := baseSuite.Tag
//	if runningTag.Get("skipSuite") == "true" {
//		t.Skipf("Skip Suite")
//	}
//	skipCaseStr := runningTag.Get("skipCase")
//	if runningTag.Get("skipCase") != "" {
//		skipCases = strings.Split(skipCaseStr, ",")
//	}
//}
//skipCases := getSkipCases(tagInfos)
