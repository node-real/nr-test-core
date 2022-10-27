package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

type DataOptTest struct {
	nrsuite.NRBaseSuite
}

func TestDataOptSuite(t *testing.T) {
	nrsuite.Run(t, new(DataOptTest))
}

func (t *DataOptTest) Test_1() {
	data := t.DataOperator.ReadCustomCaseData("./data/custom_data")
	fmt.Println(data)
}
