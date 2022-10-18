package test

import (
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

// Tags:: level:p1  net:eth,abc
type HttpTest struct {
	nrsuite.NRBaseSuite
}

func TestHttp(t *testing.T) {
	nrsuite.Run(t, new(HttpTest))
}

func (t *HttpTest) Test_Http() {

	//headers := map[string]string{
	//	"Content-Type": "application/json",
	//}
	//req := http.Request{
	//	Method:  "POST",
	//	Headers: headers,
	//}
	//t.Log.Info(".......")
	//res, err := t.Http.Call(req)
	//t.NoError(err)
	//result := t.Checker.CheckJsonValue("", res.Body)
	//t.True(result, "")
}
