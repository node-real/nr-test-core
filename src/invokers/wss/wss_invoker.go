package wss

import (
	"github.com/gorilla/websocket"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/log"
	"time"
)

type WssInvoker struct {
	//Channel WssChannel
}

func (wss *WssInvoker) SendMsg(host string, msg *rpc.RpcMessage) (*http.Response, error) {
	return wss.SendMsgWithRetry(host, msg, 0)
}

func (wss *WssInvoker) SendMsgWithRetry(host string, msg *rpc.RpcMessage, retry int) (*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		if retry <= 2 {
			return wss.SendMsgWithRetry(host, msg, retry+1)
		}
		log.Error("dial:", err)
		return nil, err
	}
	defer webClient.Close()
	err = webClient.WriteJSON(msg)
	if err != nil {
		if retry <= 2 {
			return wss.SendMsgWithRetry(host, msg, retry+1)
		}
		log.Error(err)
		return nil, err
	}

	interval := 30 * time.Second
	err = webClient.SetReadDeadline(time.Now().Add(interval))
	for {
		_, message, err := webClient.ReadMessage()
		if err != nil {
			if retry <= 2 {
				return wss.SendMsgWithRetry(host, msg, retry+1)
			}
			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		return &http.Response{Body: string(message)}, nil
	}
}

func (wss *WssInvoker) SendMsgAndClose(host string, msg *rpc.RpcMessage) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		log.Error("dial:", err)
	}
	err = webClient.WriteJSON(msg)
	webClient.Close()
}

func (wss *WssInvoker) SendMsgAndCloseWhenSub(host string, msg *rpc.RpcMessage, retry int) (*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	webClientClose, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		if retry <= 2 {
			return wss.SendMsgAndCloseWhenSub(host, msg, retry+1)
		}
		log.Error("dial:", err)
		return nil, err
	}
	defer webClient.Close()
	err = webClient.WriteJSON(msg)
	if err != nil {
		if retry <= 2 {
			return wss.SendMsgAndCloseWhenSub(host, msg, retry+1)
		}
		log.Error(err)
		return nil, err
	}

	interval := 30 * time.Second
	err = webClient.SetReadDeadline(time.Now().Add(interval))
	for {
		webClientClose.Close()
		_, message, err := webClient.ReadMessage()
		if err != nil {
			if retry <= 2 {
				return wss.SendMsgAndCloseWhenSub(host, msg, retry+1)
			}
			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		return &http.Response{Body: string(message)}, nil
	}
}

func (wss *WssInvoker) SendBatchMsg(host string, req []*rpc.RpcMessage, retry int) (*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		if retry <= 2 {
			return wss.SendBatchMsg(host, req, retry+1)
		}
		return nil, err
	}
	defer webClient.Close()

	err = webClient.WriteJSON(req)
	if err != nil {
		if retry <= 2 {
			return wss.SendBatchMsg(host, req, retry+1)
		}
		log.Error(err)
		return nil, err
	}

	interval := 30 * time.Second
	err = webClient.SetReadDeadline(time.Now().Add(interval))
	for {
		_, message, err := webClient.ReadMessage()
		if err != nil {
			if retry <= 2 {
				return wss.SendBatchMsg(host, req, retry+1)
			}
			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		return &http.Response{Body: string(message)}, nil
	}
}

// not subscribe
func (wss *WssInvoker) SendParallel(host string, req []*rpc.RpcMessage) ([]*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		return nil, err
	}
	defer webClient.Close()

	go func() {
		for _, r := range req {
			webClient.WriteJSON(r)
		}
	}()

	count := 0
	res := make([]*http.Response, len(req))
	for {
		_, message, err := webClient.ReadMessage()
		if err != nil {

			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		res[count] = &http.Response{Body: string(message)}
		count++
		if count == len(req) {
			return res, nil
		}

	}
}

func (wss *WssInvoker) SendParallelWithBatch(host string, req []*rpc.RpcMessage, batch [][]*rpc.RpcMessage) ([]*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		return nil, err
	}
	defer webClient.Close()

	go func() {
		for _, r := range req {
			webClient.WriteJSON(r)
		}
		for _, b := range batch {
			webClient.WriteJSON(b)
		}
	}()

	count := 0
	res := make([]*http.Response, len(req)+len(batch))
	for {
		_, message, err := webClient.ReadMessage()
		if err != nil {

			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		res[count] = &http.Response{Body: string(message)}
		count++
		if count == len(req)+len(batch) {
			return res, nil
		}

	}
}

func (wss *WssInvoker) GetMsg(host string, msg *rpc.RpcMessage, count int) ([]string, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer webClient.Close()

	err = webClient.WriteJSON(msg)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	result := []string{}

	interval := 30 * time.Second
	err = webClient.SetReadDeadline(time.Now().Add(interval))
	for i := 0; i < count; i++ {
		_, message, err := webClient.ReadMessage()
		if err != nil {
			log.Error("read:", err)
			return result, err
		}
		log.Debugf("Received: %s.\n", message)
		result = append(result, string(message))

	}
	return result, err
}

func (wss *WssInvoker) GetMsgWithTimeout(host string, msg *rpc.RpcMessage, count int, timeout int) ([]string, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		log.Fatal("GetIntervalWsMsg.dial:", err)
	}

	defer webClient.Close()

	err = webClient.WriteJSON(msg)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	result := []string{}

	// todo eg  1000 when more than  30s
	err = webClient.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	for i := 0; i < count; i++ {
		_, message, err := webClient.ReadMessage()
		if err != nil {
			log.Error("GetIntervalWsMsg.read:", err)
			return result, err
		}
		log.Debugf("Received: %s.\n", message)
		result = append(result, string(message))

	}
	return result, err
}
