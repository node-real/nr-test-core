package utils

import (
	"errors"
	"fmt"
	"github.com/node-real/nr-test-core/src/log"
	"runtime"
)

type Utils struct {
	MathUtils
	FileUtils
	StringUtils
}

var isLocal = runtime.GOOS == "darwin"

func IsLocal() bool {
	return isLocal
}

func SetIsLocal(v bool) {
	isLocal = v
}

func RunFunWithRetry(f func() error, retryCount int) error {
	var err error
	for i := 0; i <= retryCount; i++ {
		func() {
			defer func() {
				err1 := recover()
				if err1 != nil {
					err = errors.New(fmt.Sprint("panic error:", err1))
					log.Error(fmt.Sprint("panic error:", err1))
				}
			}()
			err = f()
		}()
		if err == nil {
			break
		}
	}
	return err
}
