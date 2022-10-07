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

type AptosWrapperClient struct {
	Account *Account
	config  *BaseConfig
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

// Init Client
func InitNewClinet(url string, furl string) *AptosWrapperClient {
	client := new(AptosWrapperClient)
	config := BaseConfig{
		Url:  url,
		Furl: furl,
	}
	client.config = &config
	client.Account = InitAccount()
	return client
}

func InitSeedClinet(url string, furl string, seed string) *AptosWrapperClient {
	client := new(AptosWrapperClient)
	config := BaseConfig{
		Url:  url,
		Furl: furl,
	}
	client.config = &config
	client.Account = InitAccountWithSeed(seed)
	return client
}

// Returns the latest ledger information.
func (r *AptosWrapperClient) GetLedgerInfo() (*http.Response, error) {
	req := new(http.Request)
	req.Path = urlLedgerinformation
	req.Host = r.config.Url
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Returns the latest account core data resource.
func (r *AptosWrapperClient) GetAccount() (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccount
	req.Host = r.config.Url
	req.PathParam = map[string]string{"address": "0x" + r.Account.Address()}

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
func (r *AptosWrapperClient) GetAccountResources(version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountResources
	req.Host = r.config.Url

	req.PathParam = map[string]string{"address": r.Account.Address()}
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

func (r *AptosWrapperClient) GetAccountResourcesByType(resourceType string, version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountResourcesByType
	req.Host = r.config.Url
	req.PathParam = map[string]string{"address": r.Account.Address(), "resource_type": resourceType}
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

func (r *AptosWrapperClient) GetAccountMoudle(address string, version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountMoudle
	req.Host = r.config.Url
	if address != "" {
		req.PathParam = map[string]string{"address": address}
	} else {
		req.PathParam = map[string]string{"address": r.Account.Address()}
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

func (r *AptosWrapperClient) GetAccountMoudlesById(address string, module_name string, version string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAccountMoudleByid
	req.Host = r.config.Url
	if address != "" {
		req.PathParam = map[string]string{"address": address, "module_name": module_name}
	} else {
		req.PathParam = map[string]string{"address": r.Account.Address(), "module_name": module_name}
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

func (r *AptosWrapperClient) GetTransactions(limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetTransactions
	req.Host = r.config.Url
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

func (r *AptosWrapperClient) GetAcccountTransaction(limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetAcccountTransaction
	req.Host = r.config.Url
	req.PathParam = map[string]string{"address": r.Account.Address()}
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

func (r *AptosWrapperClient) GetEvent(enventkey string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetEvent
	req.Host = r.config.Url
	req.PathParam = map[string]string{"event_key": enventkey}
	req.QueryParam = make(map[string]string)
	req.Method = "GET"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *AptosWrapperClient) GetEventByHandle(address string, event_handle_struct string, field_name string, limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetEventByHandle
	req.Host = r.config.Url
	if address == "" {
		req.PathParam = map[string]string{"address": r.Account.Address(), "event_handle_struct": event_handle_struct, "field_name": field_name}
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

func (r *AptosWrapperClient) SubmitTransaction(body *TransactionBody) (*http.Response, error) {
	//todo build body and provider move function

	req := new(http.Request)
	req.Path = urlSubmitTransaction
	req.Host = r.config.Url
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

func (r *AptosWrapperClient) SignTransaction(body *TransactionSignBody) (*http.Response, error) {
	req := new(http.Request)
	req.Path = urlSignTransaction
	req.Host = r.config.Url
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

func (r *AptosWrapperClient) GetTransactionsByHash(hashorversion string, limit string, start string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlGetTransactionsByHash
	req.Host = r.config.Url
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

func (r *AptosWrapperClient) Charge(amount string) (*http.Response, error) {

	req := new(http.Request)
	req.Path = urlCharge
	req.Host = r.config.Furl
	req.QueryParam = map[string]string{"amount": amount, "address": r.Account.Address()}
	req.Method = "POST"
	res, err := req.CallSplitUrl()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func setVersion(version string) {
	if version != "" {

	}

}

//---------------- test tool line-------------------

// provider information for test
// return  TransactHandle
func (r *AptosWrapperClient) GetTransactionsInfo(limit string, start string) string {
	req := new(http.Request)
	req.Path = urlGetTransactions
	req.Host = r.config.Url
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

func (r *AptosWrapperClient) AccoutSequence(new bool) string {
	if new {
		res, err := r.GetAccount()
		if err != nil {
			return "err"
		}
		number := http.GetBodyParam(res, "$.sequence_number")
		if number == "err" {
			return "err"
		}
		r.config.Sequence = number
		return number
	} else {
		return r.config.Sequence
	}
}

func (r *AptosWrapperClient) Payload(reciver string, amount string) (*http.Response, error) {
	number := r.AccoutSequence(true)

	payload := buildPayload(reciver, amount)
	signbody := buildTransignBody(r.Account, payload, number)
	sginres, err := r.SignTransaction(signbody)
	if err != nil {
		return nil, err
	}
	transbody := buildTranBodyFromSign(r.Account, signbody, sginres)
	res, err := r.SubmitTransaction(transbody)
	return res, err

}

func (r *AptosWrapperClient) PayloadMoudle(path string) (*http.Response, error) {
	number := r.AccoutSequence(true)

	payload := buildModulePayload(path)
	signbody := buildTransignBody(r.Account, payload, number)
	sginres, err := r.SignTransaction(signbody)
	if err != nil {
		return nil, err
	}
	transbody := buildTranBodyFromSign(r.Account, signbody, sginres)
	res, err := r.SubmitTransaction(transbody)
	return res, err

}

func (r *AptosWrapperClient) CraeteAccount() (*http.Response, error) {
	number := "0"
	payload := buildCreateAccountPayload(r.Account.address)
	signbody := buildTransignBody(r.Account, payload, number)
	sginres, err := r.SignTransaction(signbody)
	if err != nil {
		return nil, err
	}
	transbody := buildTranBodyFromSign(r.Account, signbody, sginres)
	res, err := r.SubmitTransaction(transbody)
	return res, err

}

func (r *AptosWrapperClient) WaitForAccount() bool {

	res, err := r.GetAccount()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1) * time.Second)
		res, err = r.GetAccount()
		if err == nil && res.Code != 200 {
			break
		}
	}
	if res.Code == 200 {
		return true
	}
	return false
}

func (r *AptosWrapperClient) WaitForTransaction(hash string) bool {

	res, err := r.GetTransactionsByHash(hash, "", "")

	for i := 0; err == nil && http.GetBodyParam(res, "$.status") != "pending_transaction" && i < 10; i++ {
		time.Sleep(time.Duration(1) * time.Second)
		res, err = r.GetTransactionsByHash(hash, "", "")
	}
	if http.GetBodyParam(res, "$.status") != "pending_transaction" {
		return true
	}
	return false
}

func (r *AptosWrapperClient) ClientAddress() string {
	return r.Account.Address()
}

func (r *AptosWrapperClient) ClientSeed() string {
	return r.Account.GetSeed()
}
