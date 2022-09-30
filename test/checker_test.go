package test

import (
	"github.com/node-real/nr-test-core/src/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CheckerTest struct {
	core.NRBaseSuite
	a string
}

func TestChecker(t *testing.T) {
	suite.Run(t, new(CheckerTest))
}

func (t *CheckerTest) Test_DiffJson() {
	json1 := "{\n    \"data\": {\n        \"pairs\": [\n            {\n                \"id\": \"0x02359703154967eec7406c59c8ffc33fed0293f5\",\n                \"reserve0\": \"0\",\n                \"reserve1\": \"0\",\n                \"reserveUSD\": \"0\",\n                \"token0\": {\n                    \"id\": \"0x230c5c04f7ba9ae043bec002b7ed41b2d5df8a5f\",\n                    \"name\": \"NRC\",\n                    \"symbol\": \"NRC\"\n                },\n                \"token0Price\": \"0\",\n                \"token1\": {\n                    \"id\": \"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c\",\n                    \"name\": \"Wrapped BNB\",\n                    \"symbol\": \"WBNB\"\n                },\n                \"token1Price\": \"0\",\n                \"volumeUSD\": \"0\"\n            }\n        ]\n    }\n}"
	json2 := "{\n    \"data\": {\n        \"pairs\": []\n    }\n}"
	diffs0 := map[string][]interface{}{}
	t.Checker.DiffJson(json1, json2, diffs0)
	t.Assertions.Equal(1, len(diffs0))
}
