package core

import (
	"github.com/stretchr/testify/suite"
)

type NRBaseSuite struct {
	suite.Suite
	TestDriver
}

func (baseSuite *NRBaseSuite) SetupTestSuite() {

}

func (baseSuite *NRBaseSuite) TearDownTestSuite() {

}

func (baseSuite *NRBaseSuite) BeforeTest() {

}

func (baseSuite *NRBaseSuite) AfterTest() {
	// 收集测试日志
	// 收集测试报告
}
