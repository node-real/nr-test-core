package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

// Tags:: level:p0 net:eth,abc
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

func (t *DataOptTest) Test_2() {
	data := t.DataOperator.GetSecretData("/qa/testplatform/wbnb_api_key")
	t.Assertions.Equal(data, "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
}
