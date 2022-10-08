package aptos

import (
	"encoding/json"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"time"
)

type BaseConfig struct {
	//rest-server
	Url string
	//Faucet endpoint:
	Furl string
	//the local cache of account Sequence
	Sequence string
}

const (
	urlCharge                    = "/mint"
	urlLedgerinformation         = "/v1"
	urlGetAccount                = "/v1/accounts/{address}"
	urlGetAccountResources       = "/v1/accounts/{address}/resources"
	urlGetAccountResourcesByType = "/v1/accounts/{address}/resource/{resource_type}"

	urlGetAccountMoudle       = "/v1/accounts/{address}/modules"
	urlGetAccountMoudleByid   = "/v1/accounts/{address}/module/{module_name}"
	urlGetTransactions        = "/v1/transactions"
	urlSubmitTransaction      = "/v1/transactions"
	urlGetAcccountTransaction = "/v1/accounts/{address}/transactions"
	urlGetTransactionsByHash  = "/v1/transactions/{txn_hash_or_version}"
	urlSignTransaction        = "/v1/transactions/encode_submission"
	urlGetEvent               = "/v1/events/{event_key}"
	urlGetEventByHandle       = "/v1/accounts/{address}/events/{event_handle_struct}/{field_name}"
	//todo
	//move module add and todo
	urlTableItemByHandle = "/v1/tables/{table_handle}/items"
)

type AptosClientWrapper struct {
	Account *Account
	config  *BaseConfig
}

func (c *AptosClientWrapper) InitClient(url string, furl string) *AptosClientWrapper {
	//client := new(AptosClientWrapper)
	config := BaseConfig{
		Url:  url,
		Furl: furl,
	}
	c.config = &config
	c.Account = InitAccount()
	return c
}

func (c *AptosClientWrapper) InitSeedClient(url string, furl string, seed string) *AptosClientWrapper {
	config := BaseConfig{
		Url:  url,
		Furl: furl,
	}
	c.config = &config
	c.Account = InitAccountWithSeed(seed)
	return c
}

// Returns the latest ledger information.
func (c *AptosClientWrapper) GetLedgerInfo() (*http.Response, error) {
	req := new(http.Request)
	req.Path = urlLedgerinformation
	req.Host = c.config.Url
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Returns the latest account core data resource.
func (c *AptosClientWrapper) GetAccount() (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccount
	req.Host = c.config.Url
	req.PathParam = map[string]string{"address": "0x" + c.Account.Address()}

	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// This API returns account resources for a specific ledger version (AKA transaction version). If not present, the latest version is used
// version:The version of the latest transaction in the ledger.
// version default:""/lastversion
func (c *AptosClientWrapper) GetAccountResources(version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountResources
	req.Host = c.config.Url

	req.PathParam = map[string]string{"address": c.Account.Address()}
	if version != "" {
		req.QueryParam = map[string]string{"version": version}
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetAccountResourcesByType(resourceType string, version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountResourcesByType
	req.Host = c.config.Url
	req.PathParam = map[string]string{"address": c.Account.Address(), "resource_type": resourceType}
	if version != "" {
		req.QueryParam = map[string]string{"version": version}
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetAccountMoudle(address string, version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountMoudle
	req.Host = c.config.Url
	if address != "" {
		req.PathParam = map[string]string{"address": address}
	} else {
		req.PathParam = map[string]string{"address": c.Account.Address()}
	}
	if version != "" {
		req.QueryParam = map[string]string{"version": version}
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetAccountMoudlesById(address string, module_name string, version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountMoudleByid
	req.Host = c.config.Url
	if address != "" {
		req.PathParam = map[string]string{"address": address, "module_name": module_name}
	} else {
		req.PathParam = map[string]string{"address": c.Account.Address(), "module_name": module_name}
	}
	if version != "" {
		req.QueryParam = map[string]string{"version": version}
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetTransactions(limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetTransactions
	req.Host = c.config.Url
	req.QueryParam = make(map[string]string)
	if limit != "" {
		req.QueryParam["limit"] = limit
	}
	if start != "" {
		req.QueryParam["start"] = start
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetAcccountTransaction(limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAcccountTransaction
	req.Host = c.config.Url
	req.PathParam = map[string]string{"address": c.Account.Address()}
	req.QueryParam = make(map[string]string)
	if limit != "" {
		req.QueryParam["limit"] = limit
	}
	if start != "" {
		req.QueryParam["start"] = start
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetEvent(enventkey string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetEvent
	req.Host = c.config.Url
	req.PathParam = map[string]string{"event_key": enventkey}
	req.QueryParam = make(map[string]string)
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetEventByHandle(address string, event_handle_struct string, field_name string, limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetEventByHandle
	req.Host = c.config.Url
	if address == "" {
		req.PathParam = map[string]string{"address": c.Account.Address(), "event_handle_struct": event_handle_struct, "field_name": field_name}
	} else {
		req.PathParam = map[string]string{"address": address, "event_handle_struct": event_handle_struct, "field_name": field_name}
	}

	req.QueryParam = make(map[string]string)
	if limit != "" {
		req.QueryParam["limit"] = limit
	}
	if start != "" {
		req.QueryParam["start"] = start
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) SubmitTransaction(body *TransactionBody) (*http.Response, error) {
	//todo build body and provider move function

	req := new(http.Request)
	req.Path = urlSubmitTransaction
	req.Host = c.config.Url
	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req.Body = string(bytes[:])
	req.Method = "POST"
	req.Headers = map[string]string{"Content-Type": "application/json"}
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) SignTransaction(body *TransactionSignBody) (*http.Response, error) {
	req := new(http.Request)
	req.Path = urlSignTransaction
	req.Host = c.config.Url
	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req.Body = string(bytes[:])
	req.Headers = map[string]string{"Content-Type": "application/json"}

	req.Method = "POST"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) GetTransactionsByHash(hashorversion string, limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetTransactionsByHash
	req.Host = c.config.Url
	req.PathParam = map[string]string{"txn_hash_or_version": hashorversion}
	req.QueryParam = make(map[string]string)
	if limit != "" {
		req.QueryParam["limit"] = limit
	}
	if start != "" {
		req.QueryParam["start"] = start
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AptosClientWrapper) Charge(amount string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlCharge
	req.Host = c.config.Furl
	req.QueryParam = map[string]string{"amount": amount, "address": c.Account.Address()}
	req.Method = "POST"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

//func setVersion(version string) {
//	if version != "" {
//
//	}
//
//}

//---------------- test tool line-------------------

// provider information for test
// return  TransactHandle
func (c *AptosClientWrapper) GetTransactionsInfo(limit string, start string) string {
	req := new(http.Request)
	req.Path = urlGetTransactions
	req.Host = c.config.Url
	req.QueryParam = make(map[string]string)
	if limit != "" {
		req.QueryParam["limit"] = limit
	}
	if start != "" {
		req.QueryParam["start"] = start
	}
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return "error"
	}
	hash := http.GetBodyParam(res, "&.hash")
	return hash
}

func (c *AptosClientWrapper) AccoutSequence(new bool) string {
	if new {
		res, err := c.GetAccount()
		if err != nil {
			return "err"
		}
		number := http.GetBodyParam(res, "$.sequence_number")
		if number == "err" {
			return "err"
		}
		c.config.Sequence = number
		return number
	} else {
		return c.config.Sequence
	}
}

func (c *AptosClientWrapper) Payload(reciver string, amount string) (*http.Response, error) {
	number := c.AccoutSequence(true)

	payload := buildPayload(reciver, amount)
	signbody := buildTransignBody(c.Account, payload, number)
	sginres, err := c.SignTransaction(signbody)
	if err != nil {
		return nil, err
	}
	transbody := buildTranBodyFromSign(c.Account, signbody, sginres)
	res, err := c.SubmitTransaction(transbody)
	return res, err

}

func (c *AptosClientWrapper) PayloadMoudle(path string) (*http.Response, error) {
	number := c.AccoutSequence(true)

	payload := buildModulePayload(path)
	signbody := buildTransignBody(c.Account, payload, number)
	sginres, err := c.SignTransaction(signbody)
	if err != nil {
		return nil, err
	}
	transbody := buildTranBodyFromSign(c.Account, signbody, sginres)
	res, err := c.SubmitTransaction(transbody)
	return res, err

}

func (c *AptosClientWrapper) CraeteAccount() (*http.Response, error) {
	number := "0"
	payload := buildCreateAccountPayload(c.Account.address)
	signbody := buildTransignBody(c.Account, payload, number)
	sginres, err := c.SignTransaction(signbody)
	if err != nil {
		return nil, err
	}
	transbody := buildTranBodyFromSign(c.Account, signbody, sginres)
	res, err := c.SubmitTransaction(transbody)
	return res, err

}

func (c *AptosClientWrapper) WaitForAccount() bool {

	res, err := c.GetAccount()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1) * time.Second)
		res, err = c.GetAccount()
		if err == nil && res.Code != 200 {
			break
		}
	}
	if res.Code == 200 {
		return true
	}
	return false
}

func (c *AptosClientWrapper) WaitForTransaction(hash string) bool {

	res, err := c.GetTransactionsByHash(hash, "", "")

	for i := 0; err == nil && http.GetBodyParam(res, "$.status") != "pending_transaction" && i < 10; i++ {
		time.Sleep(time.Duration(1) * time.Second)
		res, err = c.GetTransactionsByHash(hash, "", "")
	}
	if http.GetBodyParam(res, "$.status") != "pending_transaction" {
		return true
	}
	return false
}

func (c *AptosClientWrapper) ClientAddress() string {
	return c.Account.Address()
}

func (c *AptosClientWrapper) ClientSeed() string {
	return c.Account.GetSeed()
}
