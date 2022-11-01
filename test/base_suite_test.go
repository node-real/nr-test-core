package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

// Tags:: level:p0
type BaseSuiteTest struct {
	nrsuite.NRBaseSuite
}

func TestSuite(t *testing.T) {
	nrsuite.Run(t, new(BaseSuiteTest))
}

// Tags:: level:p0
func (t *BaseSuiteTest) Test_1() {
	fmt.Println(t.RunningConfig.TestParams["NoderealRpcUrl"])
	fmt.Println("*****1")
}

// Tags:: level:P1
func (t *BaseSuiteTest) Test_2() {
	fmt.Println("*****2")
}

// Tags:: level:p2 net:abc
func (t *BaseSuiteTest) Test_3() {
	fmt.Println("*****2")
}

// Tags:: level:p3 net:abc
func (t *BaseSuiteTest) Test_4() {
	fmt.Println("*****2")
}

func (t *BaseSuiteTest) Test_5() {
	fmt.Println("*****2")
}
