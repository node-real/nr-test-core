package utils

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/node-real/nr-test-core/src/log"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type StringUtils struct {
}

func (s *StringUtils) GetStringInBetween(str string, start string, end string) string {
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

func (s *StringUtils) ConvertStrToInt(intStr string) (int, error) {
	i, err := strconv.Atoi(intStr)
	return i, err
}

func (s *StringUtils) HexStrToBigInt(hex string) *big.Int {
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

func (s *StringUtils) BigIntToHexString(bigInt *big.Int) string {
	if nil == bigInt {
		return "0x0"
	}
	hex_value := strconv.FormatInt(bigInt.Int64(), 16)
	return "0x" + hex_value
}

func (s *StringUtils) Int64ToHexString(intValue int64) string {
	hex_value := strconv.FormatInt(intValue, 16)
	return "0x" + hex_value
}

func (s *StringUtils) StringToBigInt(string2 string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(string2, 10)
	if nil == n {
		n = new(big.Int).SetUint64(0)
	}
	return n
}

func (s *StringUtils) StringToHexString(string2 string) string {
	big1 := s.StringToBigInt(string2)
	return s.BigIntToHexString(big1)
}

func (s *StringUtils) ToJsonString(data interface{}) string {
	dataStr, ok := data.(string)
	if ok {
		return dataStr
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return ""
	}
	return string(dataJson)
}

func (s *StringUtils) RandomString(n int) string {
	return rdm(n, bytes)
}

func (s *StringUtils) RandomStringToLower(n int) string {
	return rdm(n, bytesToLower)
}

func (s *StringUtils) RandomStringToUpper(n int) string {
	return rdm(n, bytesToUpper)
}

func (s *StringUtils) RandomEthHexKey() string {
	key, _ := crypto.GenerateKey()

	keyBytes := crypto.FromECDSA(key)
	hexkey := hexutil.Encode(keyBytes)[2:]
	return hexkey
}

func rdm(n int, data string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = data[rand.Int63()%int64(len(data))]
	}
	return string(b)
}
