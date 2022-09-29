package src

import (
	"github.com/node-real/nr-test-core/src/checker"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/invokers/wss"
	"github.com/stretchr/testify/suite"
)

type NRBaseSuite struct {
	suite.Suite
	Http    http.HttpInvoker
	Rpc     rpc.RpcInvoker
	Wss     wss.WssInvoker
	Checker checker.Checker
}
