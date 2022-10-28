package test

import (
	"github.com/node-real/nr-test-core/src/core/nrdriver"
	"github.com/node-real/nr-test-core/src/core/nrsuite"
	"github.com/node-real/nr-test-core/src/log"
	"testing"
)

type LogTest struct {
	nrsuite.NRBaseSuite
}

func TestLogSuite(t *testing.T) {
	nrsuite.Run(t, new(LogTest))
}

func (t *LogTest) Test_1() {
	t.Log.Error("*****")
	log.Error("*******")
	nrdriver.Driver().Log.Error("********")
}
