package nrdriver

import (
	"github.com/node-real/nr-test-core/src/checker"
	"github.com/node-real/nr-test-core/src/core"
	"github.com/node-real/nr-test-core/src/dataopt"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/invokers/wss"
	"github.com/node-real/nr-test-core/src/log"
	"github.com/node-real/nr-test-core/src/utils"
	"sync"
)

var DConfig DriverConfig

type TestDriver struct {
	RunningConfig *core.RunningConfig
	DataOperator  *dataopt.DataOperator
	Http          *http.HttpInvoker
	Rpc           *rpc.RpcInvoker
	Wss           *wss.WssInvoker
	Checker       *checker.Checker
	Log           *log.Logger
	Utils         *utils.Utils
	CurrTask      string
	LogLever      int
}

var (
	once   sync.Once
	driver TestDriver
)

func init() {
	if core.Config != nil {
		DConfig = DriverConfig{
			LogLevel: core.Config.LogLevel,
		}
	}
	Driver()
}

func Driver() TestDriver {
	once.Do(func() {
		driver = TestDriver{}
		driver.DataOperator = &dataopt.DataOperator{}
		driver.Http = &http.HttpInvoker{}
		driver.Rpc = &rpc.RpcInvoker{}
		driver.Wss = &wss.WssInvoker{}
		driver.Checker = &checker.Checker{}
		driver.Utils = &utils.Utils{}
		var logLevel = log.InfoLog
		if DConfig != (DriverConfig{}) && logLevel != 0 {
			logLevel = log.InfoLog
		}
		driver.Log = log.InitLog(logLevel, log.Stdout)
	})
	return driver
}
