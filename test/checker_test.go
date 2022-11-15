package test

import (
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"testing"
)

// Tags:: level:p0  net:eth,abc
type CheckerSuite struct {
	nrsuite.NRBaseSuite
}

func TestChecker(t *testing.T) {
	nrsuite.Run(t, new(CheckerSuite))
}

func (t *CheckerSuite) Test_diffJson() {
	json1 := "{\n    \"data\": {\n        \"pairs\": [\n            {\n                \"id\": \"0x02359703154967eec7406c59c8ffc33fed0293f5\",\n                \"reserve0\": \"0\",\n                \"reserve1\": \"0\",\n                \"reserveUSD\": \"0\",\n                \"token0\": {\n                    \"id\": \"0x230c5c04f7ba9ae043bec002b7ed41b2d5df8a5f\",\n                    \"name\": \"NRC\",\n                    \"symbol\": \"NRC\"\n                },\n                \"token0Price\": \"0\",\n                \"token1\": {\n                    \"id\": \"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c\",\n                    \"name\": \"Wrapped BNB\",\n                    \"symbol\": \"WBNB\"\n                },\n                \"token1Price\": \"0\",\n                \"volumeUSD\": \"0\"\n            }\n        ]\n    }\n}"
	json2 := "{\n    \"data\": {\n        \"pairs\": []\n    }\n}"
	_, diffMap, diffMap1 := t.Checker.CheckJsonKVReturnDiffMap(json1, json2)
	t.Assertions.Equal(1, len(diffMap))
	t.Assertions.Equal(0, len(diffMap1))
}

func (t *CheckerSuite) Test_2() {
	t.Checker.CheckResponseGroupContains(nil, "")
}

func (t *CheckerSuite) Test_Contains_1() {
	array := []int{1, 2, 3}
	t.Assertions.True(t.Checker.IsContains(array, 1))
}

func (t *CheckerSuite) Test_Contains_2() {

	array1 := []string{"1", "2", "3"}
	t.Assertions.True(t.Checker.IsContains(array1, "1"))
}

func (t *CheckerSuite) Test_Contains_3() {

	array2 := []string{"1", "2", "3"}
	t.Assertions.False(t.Checker.IsContains(array2, 1))
}

func (t *CheckerSuite) Test_Contains_4() {

	array3 := []float64{1.2, 3.1}
	t.Assertions.True(t.Checker.IsContains(array3, 1.2))
}
