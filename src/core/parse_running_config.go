package core

import (
	"flag"
	"fmt"
	"github.com/node-real/nr-test-core/src/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var Config *RunningConfig
var once sync.Once

func InitConfig() {
	once.Do(func() {
		configV := parseRunningConfig()
		Config = &configV
	})
}

func parseRunningConfig() RunningConfig {
	//if !flag.Parsed() {
	//	flag.Parse()
	//}
	argList := flag.Args()
	log.Info("go test args :", argList)
	rConfig := RunningConfig{}
	rConfig.TestFilters = map[string]string{}
	rConfig.TestParams = map[string]string{}

	configPath := getDefaultConfigPath()
	argsTestFilter := map[string]string{}
	for _, arg := range argList {
		if strings.Contains(arg, ".yml") {
			configPath = arg
		} else {
			r := strings.Split(arg, ":")
			if len(r) == 2 {
				argsTestFilter[r[0]] = r[1]
			}
		}
	}
	parseConfigYml(configPath, &rConfig)
	for k, v := range argsTestFilter {
		rConfig.TestFilters[k] = v
	}

	log.Info("Running config:", rConfig)
	return rConfig
}

func getDefaultConfigPath() string {
	return ""
}

func parseConfigYml(ymlPath string, runningConfig *RunningConfig) {
	log.Info("Start to parse config yml file:", ymlPath)
	if ymlPath == "" {
		log.Info("Config yml path is empty.")
		return
	}
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
	log.Info("The full path of yml file:", ymlPath)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("Can not read config yaml file:", path)
	}
	configMap := map[string]interface{}{}
	yaml.Unmarshal(fileContent, &configMap)
	for k, v := range configMap {
		if k == "LogLevel" {
			value := v.(int)
			runningConfig.LogLevel = value
		}
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
