package utils

import "fmt"

type Utils struct {
	MathUtils
	FileUtils
	StringUtils
}

func RunFunWithRetry(f func() error, retryCount int) error {
	var err error
	for i := 0; i <= retryCount; i++ {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("panic error:", err)
				}
			}()
			err = f()
			if err == nil {
				return
			}
		}()
	}
	return err
}
