package test

import (
	"github.com/node-real/nr-test-core/src/core"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HttpTest struct {
	core.NRBaseSuite
}

func TestHttp(t *testing.T) {
	suite.Run(t, new(HttpTest))
}

func (t *HttpTest) Test_Http() {
	req := http.Request{
		Path: "",
		Host: "",
	}
	t.Log.Info(".......")
	res, err := t.Http.Call(req)
	t.NoError(err)
	result := t.Checker.CheckJsonValue("", res.Body)
	t.True(result, "")
}
