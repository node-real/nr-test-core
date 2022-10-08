package clients

import "github.com/node-real/nr-test-core/src/clients/aptos"

type ClientWrappers struct {
	Aptos *aptos.AptosWrapperClient
}

//
//func InitNewClients() *ClientWrappers {
//	clients := new(ClientWrappers)
//	//clients.Aptos = aptos.InitNewClinet()
//}
