package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"github.com/node-real/nr-test-core/src/invokers/http"
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

func (t *UtilsTest) Test_2() {
	dataStr := t.Utils.ToJsonString(http.Request{
		Path:   "abc",
		Method: "GET",
		Host:   "123",
	})
	fmt.Println(dataStr)

	mapData := map[string]string{
		"Test1": "123",
		"Test2": "456",
		"Test3": "789",
	}
	dataStr = t.Utils.ToJsonString(mapData)
	fmt.Println(dataStr)

	strData := "{\"Test10\":\"123\",\"Test2\":\"456\",\"Test3\":\"789\"}"
	fmt.Println(strData)
}

func (t *UtilsTest) Test_3() {
	data := "{\"Test10\":\"123\",\"Test2\":\"456\",\"Test3\":\"789\"}"
	dataStr := t.Utils.ToJsonString(data)
	t.Assert().Equal(dataStr, data)
}
