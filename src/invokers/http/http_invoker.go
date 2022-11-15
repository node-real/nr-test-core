package http

import (
	"encoding/json"
	"fmt"
	"github.com/node-real/nr-test-core/src/utils"
	"github.com/oliveagle/jsonpath"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var util = new(utils.Utils)

type HttpInvoker struct {
}

func (httpInvoker *HttpInvoker) Call(r Request) (*Response, error) {
	return httpInvoker.CallTimeOut(r, 360*time.Second)
}

func (httpInvoker *HttpInvoker) CallTimeOut(req Request, timeout time.Duration) (*Response, error) {
	//build path
	var api string
	if req.Protocol == "" && strings.Contains(req.Host, "://") {
		api = getUrl(req.Host, req.Path, req.PathParam)
	} else {
		api = getUrl(req.Protocol+"://"+req.Host, req.Path, req.PathParam)
	}

	bodyStr0 := util.ToJsonString(req.Body)
	nr, err := http.NewRequest(req.Method, api, strings.NewReader(bodyStr0))
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

	client := &http.Client{Timeout: timeout}
	sTime := time.Now()
	resp, err := client.Do(nr)
	eTime := time.Now()
	if err != nil {
		return nil, err
	}
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

// Deprecated
// CallSplitUrl coudle set url with Protocol and call the method
func (httpInvoker *HttpInvoker) CallSplitUrl(r Request) (*Response, error) {
	r.Protocol = strings.Split(r.Host, "://")[0]
	r.Host = strings.Split(r.Host, "://")[1]
	return httpInvoker.CallTimeOut(r, 360*time.Second)
}

// Deprecated
func (httpInvoker *HttpInvoker) GetBodyParam(r *Response, jpath string) string {
	var json_data interface{}
	json.Unmarshal([]byte(r.Body), &json_data)
	res, err := jsonpath.JsonPathLookup(json_data, jpath)
	if err != nil {
		return "err"
	}
	str := fmt.Sprintf("%v", res)
	return str
}

func (httpInvoker *HttpInvoker) getUrl(url string, path string, pathparam map[string]string) string {
	if pathparam != nil && len(pathparam) > 0 {
		for key := range pathparam {
			path = strings.ReplaceAll(path, "{"+key+"}", pathparam[key])
		}
	}
	return url + path
}

func getUrl(url string, path string, pathparam map[string]string) string {
	if pathparam != nil && len(pathparam) > 0 {
		for key := range pathparam {
			path = strings.ReplaceAll(path, "{"+key+"}", pathparam[key])
		}
	}
	return url + path
}
