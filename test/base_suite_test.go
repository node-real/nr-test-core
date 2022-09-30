package test

import (
	"github.com/node-real/nr-test-core/src/core"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BaseSuiteTest struct {
	core.NRBaseSuite
	a string
}

func TestBaseSuite(t *testing.T) {
	suite.Run(t, new(BaseSuiteTest))
}

func (t *BaseSuiteTest) Test_HttpInvoker() {
	t.Http.Call(http.Request{})

}
