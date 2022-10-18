package nrsuite

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type RunningConfig struct {
	//RetryCount int
	//parallelCount int
	TestFilters map[string]string
	TestParams  map[string]string
}

func parseRunningConfig() RunningConfig {
	argList := flag.Args()
	rConfig := RunningConfig{}
	rConfig.TestFilters = map[string]string{}

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
	return rConfig
}

func parseConfigYml(ymlPath string, runningConfig *RunningConfig) {
	fileContent, err := ioutil.ReadFile(ymlPath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(fileContent, runningConfig)
}
