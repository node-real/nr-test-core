package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

// Tags:: net:eth
type TagSuite struct {
	nrsuite.NRBaseSuite
}

func TestTagSuite(t *testing.T) {
	nrsuite.Run(t, new(TagSuite))
}

// Tags:: level:p1
type Tag1Suite struct {
	nrsuite.NRBaseSuite
}

func TestTag1Suite(t *testing.T) {
	nrsuite.Run(t, new(Tag1Suite))
}

// Tags:: level:p1
func (t *TagSuite) Test_1() {
	fmt.Println(t.RunningConfig.TestParams["NoderealRpcUrl"])
	fmt.Println("*****1")
}

// Tags:: level:p0
func (t *TagSuite) Test_2() {
	fmt.Println("*****2")
}

// Tags:: level:p0
func (t *Tag1Suite) Test_1() {
	fmt.Println(t.RunningConfig.TestParams["NoderealRpcUrl"])
	fmt.Println("*****1")
}

func (t *Tag1Suite) Test_2() {
	fmt.Println("*****2")
}
