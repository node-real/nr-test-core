package core

import (
	"github.com/node-real/nr-test-core/src/checker"
	"github.com/node-real/nr-test-core/src/clients"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/invokers/wss"
	"sync"
)

type TestDriver struct {
	Http    http.HttpInvoker
	Rpc     rpc.RpcInvoker
	Wss     wss.WssInvoker
	Checker checker.Checker
	Clients clients.ClientWrappers
}

var (
	once   sync.Once
	driver *TestDriver
)

func Driver() *TestDriver {
	once.Do(func() {
		driver = &TestDriver{}
	})
	return driver
}
