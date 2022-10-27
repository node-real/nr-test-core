package nrdriver

import (
	"github.com/node-real/nr-test-core/src/checker"
	"github.com/node-real/nr-test-core/src/core"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/invokers/wss"
	"github.com/node-real/nr-test-core/src/log"
	"github.com/node-real/nr-test-core/src/report"
	"github.com/node-real/nr-test-core/src/testdata"
	"github.com/node-real/nr-test-core/src/utils"
	"sync"
)

type TestDriver struct {
	RunningConfig *core.RunningConfig
	DataOperator  *testdata.DataOperator
	Http          *http.HttpInvoker
	Rpc           *rpc.RpcInvoker
	Wss           *wss.WssInvoker
	Checker       *checker.Checker
	Log           *log.Logger
	Utils         *utils.Utils
	CurrTask      string
}

var (
	once   sync.Once
	driver TestDriver
)

func Driver() TestDriver {
	once.Do(func() {
		driver = TestDriver{}
		driver.DataOperator = &testdata.DataOperator{}
		driver.Http = &http.HttpInvoker{}
		driver.Rpc = &rpc.RpcInvoker{}
		driver.Wss = &wss.WssInvoker{}
		driver.Checker = &checker.Checker{}
		driver.Log = &log.Logger{}
		report.StartReportGenerator()
	})
	return driver
}
