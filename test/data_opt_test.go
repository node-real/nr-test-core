package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrdriver"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"github.com/node-real/nr-test-core/src/log"
	"testing"
)

type DataOptTest struct {
	nrsuite.NRBaseSuite
}

func TestDataOptSuite(t *testing.T) {
	nrsuite.Run(t, new(DataOptTest))
}

func (t *DataOptTest) Test_1() {
	t.TestDriver.Log.Error("*****")
	log.Log.Error("*******")
	nrdriver.Driver().Log.Error("********")
	data := t.DataOperator.ReadCustomCaseData("./data/custom_data")
	fmt.Println(data)
}
