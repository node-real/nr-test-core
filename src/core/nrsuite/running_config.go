package nrsuite

import (
	"flag"
	"fmt"
	"github.com/node-real/nr-test-core/src/core"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func parseRunningConfig() core.RunningConfig {
	argList := flag.Args()
	rConfig := core.RunningConfig{}
	rConfig.TestFilters = map[string]string{}
	rConfig.TestParams = map[string]string{}

	for _, arg := range argList {
		if strings.Contains(arg, ".yml") {
			parseConfigYml(arg, &rConfig)
		} else {
			r := strings.Split(arg, ":")
			if len(r) == 2 {
				rConfig.TestParams[r[0]] = r[1]
			}
		}
	}
	fmt.Println(rConfig)
	return rConfig
}

func parseConfigYml(ymlPath string, runningConfig *core.RunningConfig) {
	path, err := os.Getwd()
	for i := 0; i < 10; i++ {
		fileEnum, _ := os.ReadDir(path)
		hasMod := false
		for _, f := range fileEnum {
			if f.Name() == "go.mod" {
				hasMod = true
				break
			}
		}
		if hasMod {
			path = filepath.Join(path, ymlPath)
			break
		} else {
			path = getParentDirectory(path)
		}
	}
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	configMap := map[string]interface{}{}
	yaml.Unmarshal(fileContent, &configMap)
	for k, v := range configMap {
		if k == "TestFilters" {
			aItem, _ := v.([]interface{})
			//fmt.Println(a)
			for _, vItem := range aItem {
				bItem, _ := vItem.(map[string]interface{})
				for k1, v1 := range bItem {
					v1Str := fmt.Sprintf("%v", v1)
					runningConfig.TestFilters[k1] = v1Str
				}

			}
		}
		if k == "TestParams" {
			aItem, _ := v.([]interface{})
			//fmt.Println(a)
			for _, vItem := range aItem {
				bItem, _ := vItem.(map[string]interface{})
				for k1, v1 := range bItem {
					v1Str := fmt.Sprintf("%v", v1)
					runningConfig.TestParams[k1] = v1Str
				}

			}
		}
	}
	fmt.Println(runningConfig)
}

func getParentDirectory(dir string) string {
	return substr(dir, 0, strings.LastIndex(dir, "/"))
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
