package core

type RunningConfig struct {
	ConfigPath string `yaml:"ConfigPath"`
	LogLevel   int    `yaml:"LogLevel"`
	//RetryCount int
	//parallelCount int
	TestFilters map[string]string `yaml:"TestFilters"`
	TestParams  map[string]string `yaml:"TestParams"`
}

type RunningConfig1 struct {
	//ConfigPath string
	//LogLevel   int
	//RetryCount int
	//parallelCount int
	TestFilters map[string]string `yaml:"TestFilters"`
	TestParams  map[string]string `yaml:"TestParams"`
}

type CItem struct {
	Key   string
	Value string
}
