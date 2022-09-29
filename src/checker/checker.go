package checker

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-cmp/cmp"
	"github.com/node-real/nr-test-core/src/log"
	"github.com/tidwall/gjson"
	"strings"
)

type Checker struct {
}

func (checker *Checker) IsContains(data, sub string) bool {
	if strings.Contains(data, sub) {
		return true
	}
	return false
}

func (checker *Checker) IsContainsInArray(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// IsJson check a string is json or not
func (checker *Checker) IsJson(json string) bool {
	return gjson.Valid(json)
}

// IsRange check a json field is range or not
func (checker *Checker) IsRange(actual, exp string, interval uint64) bool {
	a, err := hexutil.DecodeUint64(gjson.Get(actual, "result").String())
	if err != nil {
		return false
	}
	b, err := hexutil.DecodeUint64(gjson.Get(exp, "result").String())
	if err != nil {
		return false
	}

	return a+interval > b && b+interval > a
}

// Equals interface
func (checker *Checker) Equals(actual, exp interface{}) bool {
	return cmp.Equal(actual, exp)
}

// CheckKey just check json key ,not compare value
func (checker *Checker) CheckKey(actual, exp string) bool {
	diffs0 := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}

	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs0 = checker.DiffList("root", exp, actual, diffs0)
		diffs1 = checker.DiffList("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs0 = checker.DiffJson(exp, actual, diffs0)
		diffs1 = checker.DiffJson(actual, exp, diffs1)
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

// CheckValue check json key and value
func (checker *Checker) CheckValue(exp, actual string) bool {
	if exp == "" || actual == "" {
		log.Error("exp or actual if nil")
		return false
	}
	diffs0 := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}
	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs0 = checker.DiffList("root", exp, actual, diffs0)
		diffs1 = checker.DiffList("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs0 = checker.DiffJson(exp, actual, diffs0)
		diffs1 = checker.DiffJson(actual, exp, diffs1)
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

func (checker *Checker) CheckKeyValueNoLog(exp, actual string) (bool, map[string][]interface{}, map[string][]interface{}) {
	if exp == "" || actual == "" {
		log.Error("exp or actual if nil")
		return false, nil, nil
	}
	diffs0 := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}
	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs0 = checker.DiffList("root", exp, actual, diffs0)
		diffs1 = checker.DiffList("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs0 = checker.DiffJson(exp, actual, diffs0)
		diffs1 = checker.DiffJson(actual, exp, diffs1)
	}
	return len(diffs0) == 0 && len(diffs1) == 0, diffs0, diffs1
}

// KeyValueOpt check json key and value ,filter some key
func (checker *Checker) KeyValueOpt(exp, actual string, opt []string) bool {
	diffs := map[string][]interface{}{}
	diffs1 := map[string][]interface{}{}
	json0 := gjson.Parse(exp)
	if json0.IsArray() {
		diffs = checker.DiffList("root", exp, actual, diffs)
		diffs1 = checker.DiffList("root", actual, exp, diffs1)
	} else if json0.IsObject() {
		diffs = checker.DiffJson(exp, actual, diffs)
		diffs1 = checker.DiffJson(actual, exp, diffs1)
	}
	for k0, v0 := range diffs {
		if checker.IsContainsInArray(opt, k0) {
			delete(diffs, k0)
			continue
		}
		log.Errorf("diffs0: %s", k0)
		log.Errorf("diffs0: %v", v0)
	}
	for k0, v0 := range diffs1 {
		if checker.IsContainsInArray(opt, k0) {
			delete(diffs1, k0)
			continue
		}
		log.Errorf("diffs1: %s", k0)
		log.Errorf("diffs1: %v", v0)
	}
	return len(diffs) == 0 && len(diffs1) == 0
}

func (checker *Checker) DiffJson(jstr0, jstr1 string, diffs map[string][]interface{}) map[string][]interface{} {
	json0 := gjson.Parse(jstr0).Map()
	json1 := gjson.Parse(jstr1).Map()
	for k0, v0 := range json0 {
		if _, ok := json1[k0]; !ok {
			diffs[k0] = []interface{}{v0}
			continue
		}
		if v0.IsObject() {
			diffs = checker.DiffJson(v0.String(), json1[k0].String(), diffs)
		} else if v0.IsArray() {
			diffs = checker.DiffList(k0, v0.String(), json1[k0].String(), diffs)
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

func (checker *Checker) DiffList(key, jstr0, jstr1 string, diffs map[string][]interface{}) map[string][]interface{} {
	json0 := gjson.Parse(jstr0).Array()
	json1 := gjson.Parse(jstr1).Array()
	for k0, v0 := range json0 {
		if len(json1) <= k0 {
			log.Debugf("=============key:%s index:%d================", key, k0)
			log.Debugf("value0: %v", v0)
			log.Debugf("value1 is null")
			diffs[fmt.Sprintf("%s:%d", key, k0)] = []interface{}{v0}
		} else if v0.Type.String() == "JSON" {
			diffs = checker.DiffJson(v0.String(), json1[k0].String(), diffs)
		} else if v0.IsArray() {
			diffs = checker.DiffList(fmt.Sprintf("%s:%d", key, k0), v0.String(), json1[k0].String(), diffs)
		} else if json1[k0].Raw != v0.Raw {
			log.Debugf("=============key:%s index:%d================", key, k0)
			log.Debugf("value0: %v", v0)
			log.Debugf("value1: %v", json1[k0])
			diffs[fmt.Sprintf("%s:%d", key, k0)] = []interface{}{v0, json1[k0]}
		}
	}
	return diffs
}

//func GroupContain(res []*httpUtil.Response, except ...string) bool {
//	for _, e := range except {
//		result := false
//		for _, r := range res {
//			temp, _, _ := KeyValueNoLog(r.Body, e)
//			if temp {
//				result = true
//				break
//			}
//		}
//		if !result {
//			return result
//		}
//	}
//	return true
//
//}
