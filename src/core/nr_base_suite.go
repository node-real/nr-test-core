package core

import (
	"github.com/stretchr/testify/suite"
)

type NRBaseSuite struct {
	suite.Suite
	TestDriver
}

func (baseSuite *NRBaseSuite) SetupTestSuite() {
	baseSuite.Clients = Driver().Clients
	baseSuite.Http = Driver().Http
	baseSuite.Rpc = Driver().Rpc
	baseSuite.Wss = Driver().Wss
	baseSuite.Checker = Driver().Checker
	baseSuite.Clients = Driver().Clients
	baseSuite.Log = Driver().Log
}

func (baseSuite *NRBaseSuite) TearDownTestSuite() {

}

func (baseSuite *NRBaseSuite) BeforeTest() {

}

func (baseSuite *NRBaseSuite) AfterTest() {
	// 收集测试日志
	// 收集测试报告
}
