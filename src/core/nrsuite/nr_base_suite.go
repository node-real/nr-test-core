package testsuite

import (
	"github.com/node-real/nr-test-core/src/core/testdriver"
	"github.com/stretchr/testify/suite"
	"testing"
)

func Run(t *testing.T, suite NRBaseSuite) {

}

type NRBaseSuite struct {
	suite.Suite
	testdriver.TestDriver
}

func (baseSuite *NRBaseSuite) SetupTestSuite() {
}

func (baseSuite *NRBaseSuite) TearDownTestSuite() {
	//TODO: Robert
}

func (baseSuite *NRBaseSuite) BeforeTest() {
	//TODO: Robert
}

func (baseSuite *NRBaseSuite) caseFilter() {

}

//// Run takes a testing suite and runs all of the tests attached
//// to it.
//func Run(t *testing.T, suite TestingSuite) {
//	defer recoverAndFailOnPanic(t)
//
//	suite.SetT(t)
//
//	var suiteSetupDone bool
//
//	var stats *SuiteInformation
//	if _, ok := suite.(WithStats); ok {
//		stats = newSuiteInformation()
//	}
//
//	tests := []testing.InternalTest{}
//	methodFinder := reflect.TypeOf(suite)
//	suiteName := methodFinder.Elem().Name()
//
//	for i := 0; i < methodFinder.NumMethod(); i++ {
//		method := methodFinder.Method(i)
//
//		ok, err := methodFilter(method.Name)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "testify: invalid regexp for -m: %s\n", err)
//			os.Exit(1)
//		}
//
//		if !ok {
//			continue
//		}
//
//		if !suiteSetupDone {
//			if stats != nil {
//				stats.Start = time.Now()
//			}
//
//			if setupAllSuite, ok := suite.(SetupAllSuite); ok {
//				setupAllSuite.SetupSuite()
//			}
//
//			suiteSetupDone = true
//		}
//
//		test := testing.InternalTest{
//			Name: method.Name,
//			F: func(t *testing.T) {
//				parentT := suite.T()
//				suite.SetT(t)
//				defer recoverAndFailOnPanic(t)
//				defer func() {
//					r := recover()
//
//					if stats != nil {
//						passed := !t.Failed() && r == nil
//						stats.end(method.Name, passed)
//					}
//
//					if afterTestSuite, ok := suite.(AfterTest); ok {
//						afterTestSuite.AfterTest(suiteName, method.Name)
//					}
//
//					if tearDownTestSuite, ok := suite.(TearDownTestSuite); ok {
//						tearDownTestSuite.TearDownTest()
//					}
//
//					suite.SetT(parentT)
//					failOnPanic(t, r)
//				}()
//
//				if setupTestSuite, ok := suite.(SetupTestSuite); ok {
//					setupTestSuite.SetupTest()
//				}
//				if beforeTestSuite, ok := suite.(BeforeTest); ok {
//					beforeTestSuite.BeforeTest(methodFinder.Elem().Name(), method.Name)
//				}
//
//				if stats != nil {
//					stats.start(method.Name)
//				}
//
//				method.Func.Call([]reflect.Value{reflect.ValueOf(suite)})
//			},
//		}
//		tests = append(tests, test)
//	}
//	if suiteSetupDone {
//		defer func() {
//			if tearDownAllSuite, ok := suite.(TearDownAllSuite); ok {
//				tearDownAllSuite.TearDownSuite()
//			}
//
//			if suiteWithStats, measureStats := suite.(WithStats); measureStats {
//				stats.End = time.Now()
//				suiteWithStats.HandleStats(suiteName, stats)
//			}
//		}()
//	}
//
//	runTests(t, tests)
//}

func (baseSuite *NRBaseSuite) AfterTest() {
	// 收集测试日志
	// 收集测试报告
}
