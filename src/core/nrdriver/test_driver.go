package nrdriver

import (
	"github.com/node-real/nr-test-core/src/checker"
	"github.com/node-real/nr-test-core/src/core"
	"github.com/node-real/nr-test-core/src/data"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/invokers/wss"
	"github.com/node-real/nr-test-core/src/log"
	"github.com/node-real/nr-test-core/src/report"
	"github.com/node-real/nr-test-core/src/utils"
	"os"
	"sync"
)

var DConfig DriverConfig

type TestDriver struct {
	RunningConfig *core.RunningConfig
	DataOperator  *data.DataOperator
	Http          *http.HttpInvoker
	Rpc           *rpc.RpcInvoker
	Wss           *wss.WssInvoker
	Checker       *checker.Checker
	Log           *log.Logger
	Utils         *utils.Utils
	Report        *report.ReportOperator
	CurrTask      string
	LogLever      int
	ConfigPath    string
	Regain        string
}

var (
	once   sync.Once
	driver TestDriver
)

func initDriver() {
	if core.Config != nil {
		DConfig = DriverConfig{
			LogLevel: core.Config.LogLevel,
		}
	}
}

func Driver() *TestDriver {
	once.Do(func() {
		driver = TestDriver{}
		driver.DataOperator = &data.DataOperator{}
		driver.Http = &http.HttpInvoker{}
		driver.Rpc = &rpc.RpcInvoker{}
		driver.Wss = &wss.WssInvoker{}
		driver.Checker = &checker.Checker{}
		driver.Utils = &utils.Utils{}
		driver.Regain = os.Getenv("AWS_REGION")
		core.InitConfig()
		driver.RunningConfig = core.Config
		initDriver()
		var logLevel = log.InfoLog
		if DConfig != (DriverConfig{}) && logLevel != 0 {
			logLevel = log.InfoLog
		}
		log.Log.SetDebugLevel(logLevel)
		driver.Log = log.Log
	})
	return &driver
}
