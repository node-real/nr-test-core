package http

import (
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
	"strings"
)

func GetBodyParam(r *Response, jpath string) string {
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
