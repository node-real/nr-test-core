package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

// Tags:: level:P0
type TagSuite struct {
	nrsuite.NRBaseSuite
}

func TestTagSuite(t *testing.T) {
	nrsuite.Run(t, new(TagSuite))
}

// Tags:: level:P1
type Tag1Suite struct {
	nrsuite.NRBaseSuite
}

func TestTag1Suite(t *testing.T) {
	nrsuite.Run(t, new(Tag1Suite))
}

// Tags:: level:P1 skip:true
func (t *TagSuite) Test_1() {
	fmt.Println(t.RunningConfig.TestParams["NoderealRpcUrl"])
	fmt.Println("*****1")
}

// Tags:: level:P1
func (t *TagSuite) Test_2() {
	fmt.Println("*****2")
}

func (t *Tag1Suite) Test_1() {
	fmt.Println(t.RunningConfig.TestParams["NoderealRpcUrl"])
	fmt.Println("*****1")
}

// Tags:: level:P1
func (t *Tag1Suite) Test_2() {
	fmt.Println("*****2")
}
