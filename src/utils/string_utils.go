package utils

import (
	"encoding/json"
	"github.com/node-real/nr-test-core/src/log"
	"math/big"
	"strconv"
	"strings"
)

type StringUtils struct {
}

func (utils *StringUtils) GetStringInBetween(str string, start string, end string) (result string) {
	posFirst := strings.Index(str, start)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(str, end)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(start)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return str[posFirstAdjusted:posLast]
}

func (utils *StringUtils) ConvertStrToInt(intStr string) (int, error) {
	i, err := strconv.Atoi(intStr)
	return i, err
}

func (utils *StringUtils) HexStrToBigInt(hex string) *big.Int {
	n := new(big.Int)
	// handler 0x
	if "0x" == hex {
		n = new(big.Int).SetUint64(0)
		return n
	}
	if 128 == len(hex) {
		n, _ = n.SetString(hex[2:66], 16)
	} else {
		n, _ = n.SetString(hex[2:], 16)
	}
	return n
}

func (utils *StringUtils) BigIntToHexString(bigInt *big.Int) string {
	if nil == bigInt {
		return "0x0"
	}
	hex_value := strconv.FormatInt(bigInt.Int64(), 16)
	return "0x" + hex_value
}

func (utils *StringUtils) Int64ToHexString(intValue int64) string {
	hex_value := strconv.FormatInt(intValue, 16)
	return "0x" + hex_value
}

func (utils *StringUtils) StringToBigInt(string2 string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(string2, 10)
	if nil == n {
		n = new(big.Int).SetUint64(0)
	}
	return n
}

func (utils *StringUtils) StringToHexString(string2 string) string {
	big1 := utils.StringToBigInt(string2)
	return utils.BigIntToHexString(big1)
}

func (utils *StringUtils) ToJsonString(data interface{}) string {
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return ""
	}
	return string(dataJson)
}
