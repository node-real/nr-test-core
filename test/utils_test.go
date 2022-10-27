package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

type UtilsTest struct {
	nrsuite.NRBaseSuite
}

func TestUtils(t *testing.T) {
	nrsuite.Run(t, new(UtilsTest))
}

func (t *UtilsTest) Test_1() {
	result := t.Utils.StringUtils.StringToHexString("23452345")
	fmt.Println(result)
}
