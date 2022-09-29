package wss

import (
	"github.com/gorilla/websocket"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/invokers/rpc"
	"github.com/node-real/nr-test-core/src/log"
	"time"
)

func SendMsg(host string, msg *rpc.JsonRpcMessage, retry int) (*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		if retry <= 2 {
			return SendMsg(host, msg, retry+1)
		}
		log.Error("dial:", err)
		return nil, err
	}
	defer webClient.Close()
	err = webClient.WriteJSON(msg)
	if err != nil {
		if retry <= 2 {
			return SendMsg(host, msg, retry+1)
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
				return SendMsg(host, msg, retry+1)
			}
			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		return &http.Response{Body: string(message)}, nil
	}
}

func SendMsgAndClose(host string, msg *rpc.JsonRpcMessage) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		log.Error("dial:", err)
	}
	err = webClient.WriteJSON(msg)
	webClient.Close()
}

func SendMsgAndCloseWhenSub(host string, msg *rpc.JsonRpcMessage, retry int) (*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	webClientClose, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		if retry <= 2 {
			return SendMsgAndCloseWhenSub(host, msg, retry+1)
		}
		log.Error("dial:", err)
		return nil, err
	}
	defer webClient.Close()
	err = webClient.WriteJSON(msg)
	if err != nil {
		if retry <= 2 {
			return SendMsgAndCloseWhenSub(host, msg, retry+1)
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
				return SendMsgAndCloseWhenSub(host, msg, retry+1)
			}
			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		return &http.Response{Body: string(message)}, nil
	}
}

func GetMsg(host string, msg *rpc.JsonRpcMessage, count int) ([]string, error) {
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

func SendBatchMsg(host string, req []*rpc.JsonRpcMessage, retry int) (*http.Response, error) {
	webClient, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		if retry <= 2 {
			return SendBatchMsg(host, req, retry+1)
		}
		return nil, err
	}
	defer webClient.Close()

	err = webClient.WriteJSON(req)
	if err != nil {
		if retry <= 2 {
			return SendBatchMsg(host, req, retry+1)
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
				return SendBatchMsg(host, req, retry+1)
			}
			log.Error("read:", err)
			return nil, err
		}
		log.Debugf("Received: %s.\n", message)
		return &http.Response{Body: string(message)}, nil
	}
}

// not subscribe
func SendParallel(host string, req []*rpc.JsonRpcMessage) ([]*http.Response, error) {
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

func SendParallelWithBatch(host string, req []*rpc.JsonRpcMessage, batch [][]*rpc.JsonRpcMessage) ([]*http.Response, error) {
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

type WssChannel struct {
	Wsc chan string
}

func (s *WssChannel) BeforeTest(loop func()) {
	s.NewInstance()
	s.Product(loop)
	s.Garbage()
}

func (s *WssChannel) NewInstance() {
	s.Wsc = make(chan string, 1)
}

func (s *WssChannel) Product(loop func()) {
	go func() {
		loop()
	}()
}

// throw signal to avoid block
func (s *WssChannel) Garbage() {
	//thrown no receiver message
	go func() {
		for {
			select {
			case <-s.Wsc:
			default:
			}
		}
	}()
}

func (s *WssChannel) GetOneSignal() string {
	//thrown no receiver message
	var j string
	for i := range s.Wsc {
		j = i
		break
	}
	return j
}

func (s *WssChannel) GetNSignal(n int) string {
	//thrown no receiver message
	var j string
	for i := range s.Wsc {
		j = i
		n--
		if n <= 0 {
			break
		}
	}
	return j
}
