package rpc

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"strings"
)

var httpIn = http.HttpInvoker{}

type RpcInvoker struct {
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (rpcInvoker *RpcInvoker) NewMsg(method string, params ...interface{}) (*RpcMessage, error) {
	msg := &RpcMessage{Version: "2.0", ID: []byte("1"), Method: method}
	if params != nil { // prevent sending "params" as null value
		var err error
		if msg.Params, err = json.Marshal(params); err != nil {
			return nil, err
		}
	}
	return msg, nil
}

// SendMsg call  http util
func (rpcInvoker *RpcInvoker) SendMsg(host string, msg *RpcMessage) (*http.Response, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	url := strings.Split(host, "://")
	req := http.Request{
		"POST",
		url[0],
		url[1],
		"",
		string(body),
		headers,
		nil,
		nil,
		"",
	}
	res, err := httpIn.Call(req)

	if err != nil {
		return nil, err
	}
	return res, nil
}

// SendBatchMsg call  http util
func (rpcInvoker *RpcInvoker) SendBatchMsg(host string, msg []*RpcMessage) (*http.Response, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	//
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	url := strings.Split(host, "://")
	req := http.Request{
		"POST",
		url[0],
		url[1],
		"",
		string(body),
		headers,
		nil,
		nil,
		"",
	}
	res, err := httpIn.Call(req)
	if err != nil {
		return nil, err
	}
	//
	return res, nil
}

func (rpcInvoker *RpcInvoker) ToMsgParams(msg ethereum.CallMsg) interface{} {
	params := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		params["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		params["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		params["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		params["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return params
}

func (rpcInvoker *RpcInvoker) ToFilterParams(q ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{
		"address": q.Addresses,
		"topics":  q.Topics,
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
	}
	// 0 is latest
	if q.FromBlock == nil {
		arg["fromBlock"] = nil
	} else if q.FromBlock.Uint64() == 0 {
		arg["fromBlock"] = "latest"
	} else {
		arg["fromBlock"] = hexutil.EncodeBig(q.FromBlock)
	}
	if q.ToBlock == nil {
		arg["toBlock"] = nil
	} else if q.ToBlock.Uint64() == 0 {
		arg["toBlock"] = "latest"
	} else {
		arg["toBlock"] = hexutil.EncodeBig(q.ToBlock)
	}
	return arg, nil
}

func (rpcInvoker *RpcInvoker) ToFilterParamsOptional(q ethereum.FilterQuery) (interface{}, error) {
	//0-latest
	//-1-pending
	//-2-earliest
	arg := map[string]interface{}{}
	if q.Addresses != nil {
		arg["address"] = q.Addresses
	}
	if q.Topics != nil {
		arg["topics"] = q.Topics
	}
	//
	//if q.BlockHash != nil {
	//	arg["blockHash"] = *q.BlockHash
	//}

	// 0 is latest
	if q.FromBlock == nil || q.FromBlock.Int64() < -2 {
	} else if q.FromBlock.Int64() == 0 {
		arg["fromBlock"] = "latest"
	} else if q.FromBlock.Int64() == -1 {
		arg["fromBlock"] = "pending"
	} else if q.FromBlock.Int64() == -2 {
		arg["fromBlock"] = "earliest"
	} else {
		arg["fromBlock"] = hexutil.EncodeBig(q.FromBlock)
	}

	if q.ToBlock == nil {
	} else if q.ToBlock.Int64() == 0 {
		arg["toBlock"] = "latest"
	} else if q.ToBlock.Int64() == -1 {
		arg["toBlock"] = "pending"
	} else if q.ToBlock.Int64() == -2 {
		arg["toBlock"] = "earliest"
	} else {
		arg["toBlock"] = hexutil.EncodeBig(q.ToBlock)
	}
	return arg, nil
}
