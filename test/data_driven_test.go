package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"github.com/node-real/nr-test-core/src/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

// Tags:: level:p0
type DataSuiteTest struct {
	nrsuite.NRBaseSuite
}

func TestDataSuite(t *testing.T) {
	nrsuite.Run(t, new(DataSuiteTest))
}

func (t *DataSuiteTest) SetupSuite() {
	t.TestData = map[string][]string{}

	t.TestData["EthTop200"] = []string{
		"12113241234123",
		"abcasdfasdfasd",
		"13241234123412",
		"12113241234123",
		"abcasdfasdfasd",
		"13241234123412",
		"134123412342314",
	}
	t.TestData["Tokens"] = t.DataOperator.ReadFileLines("./data/token.log")
}

// Tags:: $RunDataKey:EthTop200 $ParallelCount:10
func (t *DataSuiteTest) Test_1(data string, tt *testing.T) {
	tt.Log("test data:", data)
	fmt.Println()
	time.Sleep(time.Second * 10)
	assert.True(tt, true, "*****")
}

// Tags:: $RunDataKey:Tokens $ParallelCount:300
func (t *DataSuiteTest) Test_2(data string, tt *testing.T) {
	tt.Log("test data:", data)
	time.Sleep(time.Millisecond * 100)
	if strings.HasPrefix(data, "9") {
		t.AppendResultData("targetToken", data)
	}
}

func (t *DataSuiteTest) AfterTest(suiteName, testName string) {
	if t.ResultData["targetToken"] != nil {
		utils.WriteFileWithLines(t.ResultData["targetToken"], "./data/result_tokens.log")
	}
}

//
//// Tags:: $RunTimes=1000 $RunDuration=1d $RunInterval=10s $RunDataKey=""
//func (t *DataSuiteTest) Test_2(data string) {
//	fmt.Println(data)
//}
//
//// Tags:: $RunTimes=100 $RunDuration=1m $RunInterval=10s $RunDataKey=""
//func (t *DataSuiteTest) Test_3(data string) {
//	fmt.Println(data)
//}
