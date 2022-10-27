package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"github.com/node-real/nr-test-core/src/invokers/http"
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
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	res, err := t.Http.Call(http.Request{
		Host:     "meganode-portal.nodereal.io",
		Protocol: "https",
		Path:     "/api/v1/users/00000000-0000-0000-0000-000000000000/styles/component/1",
		Method:   "GET",
		Headers:  headers,
	})
	t.NoError(err)
	fmt.Println(res.Body)
	result := t.Checker.CheckJsonValue("{\"code\":20000,\"msg\":\"\",\"data\":{\"product_new_tag_end_time\":\"2200-01-01T00:00:00Z\",\"web3_api_marketplace_new_tag_end_time\":\"2200-01-01T00:00:00Z\"}}\n", res.Body)
	t.True(result, "")
}
