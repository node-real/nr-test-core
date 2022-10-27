package test

import (
	"fmt"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"github.com/tidwall/gjson"
	"testing"
)

// Tags:: level:p1  net:eth,abc
type RpcTest struct {
	nrsuite.NRBaseSuite
}

func TestRpc(t *testing.T) {
	nrsuite.Run(t, new(RpcTest))
}

func (t *RpcTest) Test_eth_blockNumber() {
	nrRpcUrl := "https://bsc-mainnet.nodereal.io/v1/6fba0fffb05e4b459a3c6c4a4ca88920"
	expectRpcUrl := "https://bsc-test.binance.org/"
	msg, _ := t.Rpc.NewMsg("eth_blockNumber")
	nrResult, err1 := t.Rpc.SendMsg(nrRpcUrl, msg)
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
