package core

type RunningConfig struct {
	//RetryCount int
	//parallelCount int
	TestFilters map[string]string `yaml:"TestFilters"`
	TestParams  map[string]string `yaml:"TestParams"`
}
