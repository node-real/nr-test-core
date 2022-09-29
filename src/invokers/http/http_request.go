package http

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Method     string
	Protocol   string
	Host       string
	Path       string
	Body       string
	Headers    map[string]string
	QueryParam map[string]string
	PathParam  map[string]string
	Check      string
}

// Call rpc refactor */
func (req *Request) Call() (*Response, error) {
	return req.CallTimeOut(360 * time.Second)
}

func (req *Request) CallTimeOut(timeout time.Duration) (*Response, error) {
	//change path
	api := getUrl(req.Protocol+"://"+req.Host, req.Path, req.PathParam)
	nr, err := http.NewRequest(req.Method, api, strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}
	if req.Headers != nil {
		for k, v := range req.Headers {
			nr.Header.Add(k, v)
		}
	}
	//add query
	q := nr.URL.Query()
	if req.QueryParam != nil && len(req.QueryParam) > 0 {
		for key := range req.QueryParam {
			q.Add(key, req.QueryParam[key])
		}
	}
	nr.URL.RawQuery = q.Encode()

	sTime := time.Now()
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(nr)
	if err != nil {
		return nil, err
	}
	eTime := time.Now()
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(bodyByte)

	res := &Response{Code: resp.StatusCode, Body: bodyStr, Time: eTime.Sub(sTime).Milliseconds(), Header: resp.Header}
	//log.Debug(api, res.Time)
	return res, nil
}

// CallSplitUrl coudle set url with Protocol and call the method
func (req *Request) CallSplitUrl() (*Response, error) {
	req.Protocol = strings.Split(req.Host, "://")[0]
	req.Host = strings.Split(req.Host, "://")[1]
	return req.CallTimeOut(360 * time.Second)
}
