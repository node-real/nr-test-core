package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

type AdvanceSuiteTest struct {
	nrsuite.NRBaseSuite
}

func TestAdvanceSuite(t *testing.T) {
	nrsuite.Run(t, new(AdvanceSuiteTest))
}

func (t *AdvanceSuiteTest) Test_1() {
	t.RunFunWithRetry(func() error {
		a := []int{}
		fmt.Println(a[1])
		return nil
	}, 4)
}
