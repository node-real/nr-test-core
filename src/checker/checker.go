package checker

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-cmp/cmp"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/node-real/nr-test-core/src/log"
	"github.com/tidwall/gjson"
	http1 "net/http"
	"reflect"
	"strings"
)

type JsonDiff struct {
	HasDiff bool
	Result  string
}

type Checker struct {
}

// IsContainsInStr verifies that the orgStr contains subStr or not
// For example:
//
//	s.Checker.IsContains("abcd", "ab")
func (checker *Checker) IsContainsInStr(orgStr, subStr string) bool {
	if strings.Contains(orgStr, subStr) {
		return true
	}
	return false
}

// Deprecated
// IsContainsInStrArray verifies that the string items array contains item0 string or not
// For example:
//
//	array := []int{d1,2,3}
//	s.Checker.IsContainsItemsInArray(array, 1)
func (c *Checker) IsContainsInStrArray(items []string, item0 string) bool {
	for _, eachItem := range items {
		if eachItem == item0 {
			return true
		}
	}
	return false
}

// IsContains verifies that the  items (string, int, float) array contains item0 (string, int, float)  or not
// For example:
//
//	   1.
//				array := []int{1,2,3}
//				s.Checker.IsContains(array, 1)
//	   2.
//				array1 := []string{"1", "2", "3"}
//				t.Assertions.True(t.Checker.IsContains(array1, "1"))
//	   3.
//				array3 := []float64{1.2, 3.1}
//				t.Assertions.True(t.Checker.IsContains(array3, 1.2))
//	   4.
//				t.Assertions.True(t.Checker.IsContains("abc", "a"))
func (c *Checker) IsContains(items interface{}, item0 interface{}) bool {
	result := false
	item0V := reflect.TypeOf(item0)
	switch item0V.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		item0 = reflect.ValueOf(item0).Int()
		break
	case reflect.Float32, reflect.Float64:
		item0 = reflect.ValueOf(item0).Float()
		break
	case reflect.Interface:
		item0 = reflect.ValueOf(item0).Interface()
		break
	case reflect.String:
		item0 = reflect.ValueOf(item0).String()
		break
	}
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(items)
		for i := 0; i < s.Len(); i++ {
			item := s.Index(i)
			var itemValue interface{}
			switch item.Kind() {
			case reflect.String:
				itemValue = item.String()
				break
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				itemValue = item.Int()
				break
			case reflect.Float32, reflect.Float64:
				itemValue = item.Float()
				break
			case reflect.Interface:
				itemValue = item.Interface()
				break
			}
			if itemValue != nil && c.IsEquals(itemValue, item0) {
				result = true
				break
			}
		}
		break
	case reflect.String:
		s := reflect.ValueOf(items)
		result = strings.Contains(s.String(), item0.(string))
		break
	}
	return result
}

//func (c *Checker) IsContainsInIntArray(items []int, item0 int) bool {
//	for _, eachItem := range items {
//		if eachItem == item0 {
//			return true
//		}
//	}
//	return false
//}

//func (c *Checker) IsContainsSubArray(items0 []interface{}, items1 []interface{}) bool {
//	for _, eachItem := range items0 {
//		if eachItem == item {
//			return true
//		}
//	}
//	return false
//}

// IsJson check a string is json or not
// For example:
//
//	s.Checker.IsJson("{\n    \"data\": {\n        \"pairs\": []\n    }\n}")
func (c *Checker) IsJson(jsonStr string) bool {
	return gjson.Valid(jsonStr)
}

// IsEquals interface
func (c *Checker) IsEquals(expected, actual interface{}) bool {
	return cmp.Equal(expected, actual)
}

func (c *Checker) CheckNumberInterval(expected, actual uint64, interval uint64) bool {
	return actual+interval > expected && expected+interval > actual
}

// CheckNumberStrInterval check the interval of tow number less than a value
func (c *Checker) CheckNumberStrInterval(expected, actual string, interval uint64) bool {
	a, err := hexutil.DecodeUint64(actual)
	if err != nil {
		return false
	}
	b, err := hexutil.DecodeUint64(expected)
	if err != nil {
		return false
	}
	return c.CheckNumberInterval(a, b, interval)
}

// CheckJsonKey just check json key ,not compare value
func (c *Checker) CheckJsonKey(exp, actual string) bool {
	diffs0 := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}

	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs0 = c.diffJsonListItem("root", exp, actual, diffs0)
		diffs1 = c.diffJsonListItem("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs0 = c.diffJson(exp, actual, diffs0)
		diffs1 = c.diffJson(actual, exp, diffs1)
	}
	result := true
	for k0, _ := range diffs0 {
		if _, value := diffs1[k0]; !value {
			log.Errorf("diff0 key: %s is not exit", k0)
			result = false
		}
	}
	for k0, _ := range diffs1 {
		if _, value := diffs0[k0]; !value {
			log.Errorf("diff1 key: %s is not exit", k0)
			result = false
		}
	}
	return result
}

// CheckJsonValue check json key and value
func (c *Checker) CheckJsonValue(exp, actual string) bool {
	if exp == "" || actual == "" {
		log.Error("exp or actual if nil")
		return false
	}
	diffs0 := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}
	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs0 = c.diffJsonListItem("root", exp, actual, diffs0)
		diffs1 = c.diffJsonListItem("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs0 = c.diffJson(exp, actual, diffs0)
		diffs1 = c.diffJson(actual, exp, diffs1)
	}
	for k0, v0 := range diffs0 {
		log.Errorf("diffs0: %s", k0)
		log.Errorf("diffs0: %v", v0)
	}
	for k0, v0 := range diffs1 {
		log.Errorf("diffs1: %s", k0)
		log.Errorf("diffs1: %v", v0)
	}
	return len(diffs0) == 0 && len(diffs1) == 0
}

func (c *Checker) CheckJsonKVReturnDiffMap(exp, actual string) (bool, map[string][]interface{}, map[string][]interface{}) {
	if exp == "" || actual == "" {
		log.Error("exp or actual if nil")
		return false, nil, nil
	}
	diffs0 := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}
	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs0 = c.diffJsonListItem("root", exp, actual, diffs0)
		diffs1 = c.diffJsonListItem("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs0 = c.diffJson(exp, actual, diffs0)
		diffs1 = c.diffJson(actual, exp, diffs1)
	}
	return len(diffs0) == 0 && len(diffs1) == 0, diffs0, diffs1
}

func (Checker *Checker) CheckJsonKVReturnDiffStr(exp, actual string) string {
	var json1 map[string]interface{}
	var json2 map[string]interface{}
	json.Unmarshal([]byte(exp), &json1)
	json.Unmarshal([]byte(actual), &json2)
	_, diffStr := jsonCompare(json1, json2)
	return diffStr
}

// CheckJsonKVOpt check json key and value ,filter some key
func (c *Checker) CheckJsonKVOpt(exp, actual string, opt []string) bool {
	diffs := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}
	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs = c.diffJsonListItem("root", exp, actual, diffs)
		diffs1 = c.diffJsonListItem("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs = c.diffJson(exp, actual, diffs)
		diffs1 = c.diffJson(actual, exp, diffs1)
	}
	for k0, v0 := range diffs {
		if c.IsContainsInStrArray(opt, k0) {
			delete(diffs, k0)
			continue
		}
		log.Errorf("diffs0: %s", k0)
		log.Errorf("diffs0: %v", v0)
	}
	for k0, v0 := range diffs1 {
		if c.IsContainsInStrArray(opt, k0) {
			delete(diffs1, k0)
			continue
		}
		log.Errorf("diffs1: %s", k0)
		log.Errorf("diffs1: %v", v0)
	}
	return len(diffs) == 0 && len(diffs1) == 0
}

//func (c *Checker) CheckJsonKVWithFieldsMapping(exp, actual string, mapping map[string]string) bool {
//	result, diff0, diff1 := c.CheckJsonKVReturnDiffMap(exp, actual)
//	if !result {
//		for key, value := range diff0 {
//
//		}
//	} else {
//		return true
//	}
//}

func (c *Checker) CheckJsonGroupContains(jsonStrArray []string, exp ...string) bool {
	for _, e := range exp {
		result := false
		for _, json := range jsonStrArray {
			temp := c.CheckJsonValue(json, e)
			if temp {
				result = true
				break
			}
		}
		if !result {
			return result
		}
	}
	return true
}

func (c *Checker) CheckResponseGroupContains(resArray []*http.Response, exp ...string) bool {
	for _, e := range exp {
		result := false
		for _, r := range resArray {
			temp := c.CheckJsonValue(r.Body, e)
			if temp {
				result = true
				break
			}
		}
		if !result {
			return result
		}
	}
	return true
}

func (c *Checker) CheckTowHeaderKeyContain(exp, actual http1.Header) (bool, map[string][]string) {
	diffs0 := map[string][]string{}
	for k, value := range exp {
		_, ok := actual[k]
		if !ok {
			diffs0[k] = value
		}
	}
	result := true
	if len(diffs0) > 0 {
		result = false
	}
	return result, diffs0
}

func (c *Checker) CheckTowMapKeyContain(exp, actual map[string]interface{}) (bool, map[string]interface{}) {
	diffs0 := map[string]interface{}{}
	for k, value := range exp {
		_, ok := actual[k]
		if !ok {
			diffs0[k] = value
		}
	}
	result := true
	if len(diffs0) > 0 {
		result = false
	}
	return result, diffs0
}

func (c *Checker) CheckTowHeaderValueContain(exp, actual http1.Header) (bool, map[string][]string) {
	diffs0 := map[string][]string{}
	for k, value := range exp {
		actValue, ok := actual[k]
		if !ok || c.IsEquals(actValue, value) {
			diffs0[k] = actValue
		}
	}
	result := true
	if len(diffs0) > 0 {
		result = false
	}
	return result, diffs0
}

func (c *Checker) CheckTowMapValueContain(exp, actual map[string]interface{}) (bool, map[string]interface{}) {
	diffs0 := map[string]interface{}{}
	for k, value := range exp {
		actValue, ok := actual[k]
		if !ok || c.IsEquals(actValue, value) {
			diffs0[k] = actValue
		}
	}
	result := true
	if len(diffs0) > 0 {
		result = false
	}
	return result, diffs0
}

// Deprecated
func (c *Checker) DiffJsonReturnDiffMap(jsonStr1 string, jsonStr2 string) map[string][]interface{} {
	diffMap := map[string][]interface{}{}
	return c.diffJson(jsonStr1, jsonStr2, diffMap)
}

// Deprecated
func (c *Checker) DiffJsonReturnDiffStr(jsonStr1 string, jsonStr2 string) string {
	var json1 map[string]interface{}
	var json2 map[string]interface{}
	json.Unmarshal([]byte(jsonStr1), &json1)
	json.Unmarshal([]byte(jsonStr2), &json2)
	_, diffStr := jsonCompare(json1, json2)
	return diffStr
}

//func (c *Checker) DiffJsonWithPrecision(jsonStr1 string, jsonStr2 string) string {
//	//TODO: to robert
//	return ""
//}

//func (c *Checker) CheckListStrReturnDiffMap() {
//
//}

// Deprecated
func (c *Checker) DiffList(key, jstr0, jstr1 string, diffs map[string][]interface{}) map[string][]interface{} {
	return c.diffJsonListItem(key, jstr0, jstr1, diffs)
}

func (c *Checker) diffJson(jstr0, jstr1 string, diffs map[string][]interface{}) map[string][]interface{} {
	json0 := gjson.Parse(jstr0).Map()
	json1 := gjson.Parse(jstr1).Map()
	for k0, v0 := range json0 {
		if _, ok := json1[k0]; !ok {
			diffs[k0] = []interface{}{v0}
			continue
		}
		if v0.IsObject() {
			diffs = c.diffJson(v0.String(), json1[k0].String(), diffs)
		} else if v0.IsArray() {
			diffs = c.diffJsonListItem(k0, v0.String(), json1[k0].String(), diffs)
		} else if json1[k0].Raw != v0.Raw {
			log.Debugf("=============key: %v==================", k0)
			log.Debugf("value0: %v", v0)
			log.Debugf("value1: %v", json1[k0])
			if _, ok := diffs[k0]; !ok {
				diffs[k0] = []interface{}{v0, json1[k0]}
			} else {
				diffs[k0] = []interface{}{diffs[k0], []interface{}{v0, json1[k0]}}
			}
		}
	}
	return diffs
}

func (c *Checker) diffJsonListItem(key, jstr0, jstr1 string, diffs map[string][]interface{}) map[string][]interface{} {
	json0 := gjson.Parse(jstr0).Array()
	json1 := gjson.Parse(jstr1).Array()
	for k0, v0 := range json0 {
		if len(json1) <= k0 {
			log.Debugf("=============key:%s index:%d================", key, k0)
			log.Debugf("value0: %v", v0)
			log.Debugf("value1 is null")
			diffs[fmt.Sprintf("%s:%d", key, k0)] = []interface{}{v0}
		} else if v0.Type.String() == "JSON" {
			diffs = c.diffJson(v0.String(), json1[k0].String(), diffs)
		} else if v0.IsArray() {
			diffs = c.diffJsonListItem(fmt.Sprintf("%s:%d", key, k0), v0.String(), json1[k0].String(), diffs)
		} else if json1[k0].Raw != v0.Raw {
			log.Debugf("=============key:%s index:%d================", key, k0)
			log.Debugf("value0: %v", v0)
			log.Debugf("value1: %v", json1[k0])
			diffs[fmt.Sprintf("%s:%d", key, k0)] = []interface{}{v0, json1[k0]}
		}
	}
	return diffs
}

func jsonCompare(left, right map[string]interface{}) (bool, string) {
	n := -1
	diff := &JsonDiff{HasDiff: false, Result: ""}
	jsonDiffDict(left, right, 1, diff)
	if diff.HasDiff {
		if n < 0 {
			return !diff.HasDiff, diff.Result
		} else {
			return diff.HasDiff, processContext(diff.Result, n)
		}
	}
	return !diff.HasDiff, ""
}

func marshal(j interface{}) string {
	value, _ := json.Marshal(j)
	return string(value)
}

func jsonDiffDict(json1, json2 map[string]interface{}, depth int, diff *JsonDiff) {
	blank := strings.Repeat(" ", (2 * (depth - 1)))
	longBlank := strings.Repeat(" ", (2 * (depth)))
	diff.Result = diff.Result + "\n" + blank + "{"
	for key, value := range json1 {
		quotedKey := fmt.Sprintf("\"%s\"", key)
		if _, ok := json2[key]; ok {
			switch value.(type) {
			case map[string]interface{}:
				if _, ok2 := json2[key].(map[string]interface{}); !ok2 {
					diff.HasDiff = true
					diff.Result = diff.Result + "\n-" + blank + quotedKey + ": " + marshal(value) + ","
					diff.Result = diff.Result + "\n+" + blank + quotedKey + ": " + marshal(json2[key])
				} else {
					diff.Result = diff.Result + "\n" + longBlank + quotedKey + ": "
					jsonDiffDict(value.(map[string]interface{}), json2[key].(map[string]interface{}), depth+1, diff)
				}
			case []interface{}:
				diff.Result = diff.Result + "\n" + longBlank + quotedKey + ": "
				if _, ok2 := json2[key].([]interface{}); !ok2 {
					diff.HasDiff = true
					diff.Result = diff.Result + "\n-" + blank + quotedKey + ": " + marshal(value) + ","
					diff.Result = diff.Result + "\n+" + blank + quotedKey + ": " + marshal(json2[key])
				} else {
					jsonDiffList(value.([]interface{}), json2[key].([]interface{}), depth+1, diff)
				}
			default:
				if !reflect.DeepEqual(value, json2[key]) {
					diff.HasDiff = true
					diff.Result = diff.Result + "\n-" + blank + quotedKey + ": " + marshal(value) + ","
					diff.Result = diff.Result + "\n+" + blank + quotedKey + ": " + marshal(json2[key])
				} else {
					diff.Result = diff.Result + "\n" + longBlank + quotedKey + ": " + marshal(value)
				}
			}
		} else {
			diff.HasDiff = true
			diff.Result = diff.Result + "\n-" + blank + quotedKey + ": " + marshal(value)
		}
		diff.Result = diff.Result + ","
	}
	for key, value := range json2 {
		if _, ok := json1[key]; !ok {
			diff.HasDiff = true
			diff.Result = diff.Result + "\n+" + blank + "\"" + key + "\"" + ": " + marshal(value) + ","
		}
	}
	diff.Result = diff.Result + "\n" + blank + "}"
}

func jsonDiffList(json1, json2 []interface{}, depth int, diff *JsonDiff) {
	blank := strings.Repeat(" ", (2 * (depth - 1)))
	longBlank := strings.Repeat(" ", (2 * (depth)))
	diff.Result = diff.Result + "\n" + blank + "["
	size := len(json1)
	if size > len(json2) {
		size = len(json2)
	}
	for i := 0; i < size; i++ {
		switch json1[i].(type) {
		case map[string]interface{}:
			if _, ok := json2[i].(map[string]interface{}); ok {
				jsonDiffDict(json1[i].(map[string]interface{}), json2[i].(map[string]interface{}), depth+1, diff)
			} else {
				diff.HasDiff = true
				diff.Result = diff.Result + "\n-" + blank + marshal(json1[i]) + ","
				diff.Result = diff.Result + "\n+" + blank + marshal(json2[i])
			}
		case []interface{}:
			if _, ok2 := json2[i].([]interface{}); !ok2 {
				diff.HasDiff = true
				diff.Result = diff.Result + "\n-" + blank + marshal(json1[i]) + ","
				diff.Result = diff.Result + "\n+" + blank + marshal(json2[i])
			} else {
				jsonDiffList(json1[i].([]interface{}), json2[i].([]interface{}), depth+1, diff)
			}
		default:
			if !reflect.DeepEqual(json1[i], json2[i]) {
				diff.HasDiff = true
				diff.Result = diff.Result + "\n-" + blank + marshal(json1[i]) + ","
				diff.Result = diff.Result + "\n+" + blank + marshal(json2[i])
			} else {
				diff.Result = diff.Result + "\n" + longBlank + marshal(json1[i])
			}
		}
		diff.Result = diff.Result + ","
	}
	for i := size; i < len(json1); i++ {
		diff.HasDiff = true
		diff.Result = diff.Result + "\n-" + blank + marshal(json1[i])
		diff.Result = diff.Result + ","
	}
	for i := size; i < len(json2); i++ {
		diff.HasDiff = true
		diff.Result = diff.Result + "\n+" + blank + marshal(json2[i])
		diff.Result = diff.Result + ","
	}
	diff.Result = diff.Result + "\n" + blank + "]"
}

func processContext(diff string, n int) string {
	index1 := strings.Index(diff, "\n-")
	index2 := strings.Index(diff, "\n+")
	begin := 0
	end := 0
	if index1 >= 0 && index2 >= 0 {
		if index1 <= index2 {
			begin = index1
		} else {
			begin = index2
		}
	} else if index1 >= 0 {
		begin = index1
	} else if index2 >= 0 {
		begin = index2
	}
	index1 = strings.LastIndex(diff, "\n-")
	index2 = strings.LastIndex(diff, "\n+")
	if index1 >= 0 && index2 >= 0 {
		if index1 <= index2 {
			end = index2
		} else {
			end = index1
		}
	} else if index1 >= 0 {
		end = index1
	} else if index2 >= 0 {
		end = index2
	}
	pre := diff[0:begin]
	post := diff[end:]
	i := 0
	l := begin
	for i < n && l >= 0 {
		i++
		l = strings.LastIndex(pre[0:l], "\n")
	}
	r := 0
	j := 0
	for j <= n && r >= 0 {
		j++
		t := strings.Index(post[r:], "\n")
		if t >= 0 {
			r = r + t + 1
		}
	}
	if r < 0 {
		r = len(post)
	}
	return pre[l+1:] + diff[begin:end] + post[0:r+1]
}
