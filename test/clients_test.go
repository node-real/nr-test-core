package test

import (
	"github.com/node-real/nr-test-core/src/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientsTest struct {
	core.NRBaseSuite
}

func TestClients(t *testing.T) {
	suite.Run(t, new(ClientsTest))
}

//func (t *ClientsTest) Test_Apots_Client() {
//	actualAptosClient := t.Clients.Aptos.InitClient("***", "***")
//	expAptosClient := t.Clients.Aptos.InitClient("***", "***")
//	actual, err := actualAptosClient.GetAccountResources("")
//	t.NoError(err)
//	expected, err := expAptosClient.GetAccountResources("")
//	t.NoError(err)
//	res := t.Checker.KeyValueOpt(expected.Body, actual.Body, []string{"authentication_key"})
//	t.True(res, "****")
//}
