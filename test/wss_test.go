package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"github.com/tidwall/gjson"
	"testing"
)

// Tags:: level:p1  net:eth,abc
type WssSuiteTest struct {
	nrsuite.NRBaseSuite
}

func TestTest(t *testing.T) {
	nrsuite.Run(t, new(WssSuiteTest))
}

func (t *WssSuiteTest) Test_eth_blockNumber() {
	nrRpcUrl := "wss://bsc-mainnet-us.nodereal.io/ws/v1/21aa061c92c847b5b530e53adad2c1bb"
	expectRpcUrl := "https://bsc-test.binance.org/"
	msg, _ := t.Rpc.NewMsg("eth_blockNumber")
	nrResult, err1 := t.Wss.SendMsg(nrRpcUrl, msg)
	t.NoError(err1)
	exResult, err2 := t.Rpc.SendMsg(expectRpcUrl, msg)
	t.NoError(err2)
	nrNumberStr := gjson.Get(nrResult.Body, "result").String()
	exNumberStr := gjson.Get(exResult.Body, "result").String()
	result := t.Checker.CheckNumberStrInterval(nrNumberStr, exNumberStr, 10)
	fmt.Println("nr:", nrNumberStr)
	fmt.Println("ex:", exNumberStr)
	t.True(result, "")
}
