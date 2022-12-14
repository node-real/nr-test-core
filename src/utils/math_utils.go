package utils

import (
	"math/big"
	"math/rand"
	"strings"
)

const (
	bytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytesToLower = "abcdefghijklmnopqrstuvwxyz"
	bytesToUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type MathUtils struct {
}

func (m *MathUtils) IsEven(num int) bool {
	if num%2 == 0 {
		return true
	}

	return false
}

func (m *MathUtils) RandomInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return min
	}
	return rand.Intn(max-min) + min
}

func (m *MathUtils) BigIntToStringWithPrecise(bigNum *big.Int, precise uint64) string {
	result := ""
	destStr := bigNum.String()
	destLen := uint64(len(destStr))
	if precise >= destLen {
		var i uint64 = 0
		prefix := "0."
		for ; i < precise-destLen; i++ {
			prefix += "0"
		}
		result = prefix + destStr
	} else { // add "."
		pointIndex := destLen - precise
		result = destStr[0:pointIndex] + "." + destStr[pointIndex:]
	}
	result = removeZeroAtTail(result)
	return result
}

// delete no need "0" at last of result
func removeZeroAtTail(str string) string {
	i := len(str) - 1
	for ; i >= 0; i-- {
		if str[i] != '0' {
			break
		}
	}
	str = str[:i+1]
	// delete "." at last of result
	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}
	return str
}
func ToIntByUint16(num uint16) *big.Int {
	return big.NewInt(int64(num))
}
func ToIntByPrecise(str string, precise uint64) *big.Int {
	result := new(big.Int)
	splits := strings.Split(str, ".")
	if len(splits) == 1 { // doesn't contain "."
		var i uint64 = 0
		for ; i < precise; i++ {
			str += "0"
		}
		intValue, ok := new(big.Int).SetString(str, 10)
		if ok {
			result.Set(intValue)
		}
	} else if len(splits) == 2 {
		value := new(big.Int)
		ok := false
		floatLen := uint64(len(splits[1]))
		if floatLen <= precise { // add "0" at last of str
			parseString := strings.Replace(str, ".", "", 1)
			var i uint64 = 0
			for ; i < precise-floatLen; i++ {
				parseString += "0"
			}
			value, ok = value.SetString(parseString, 10)
		} else { // remove redundant digits after "."
			splits[1] = splits[1][:precise]
			parseString := splits[0] + splits[1]
			value, ok = value.SetString(parseString, 10)
		}
		if ok {
			result.Set(value)
		}
	}

	return result
}

func Div(x *big.Int, y *big.Int) *big.Int {
	return big.NewInt(0).Div(x, y)
}
func Mul(x *big.Int, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func Sub(x *big.Int, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func Add(x *big.Int, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

func GetExpUtilization(borrow, cash, reserves *big.Int) *big.Int {
	x := Sub(Add(borrow, cash), reserves)
	y := Mul(borrow, ToIntByPrecise("1", 18))
	return Div(y, x)
}
func CalBorrowRate(utils, kink, jump, normal, base float64) float64 {
	if utils <= kink {
		return utils*normal + base
	} else {
		return kink*normal + base + (utils-kink)*jump
	}
}
func Div2float(x *big.Int, y *big.Int) *big.Float {
	if x.Cmp(big.NewInt(0)) == 0 || y.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(1)
	}
	k := big.NewFloat(0).SetInt(x)
	h := big.NewFloat(0).SetInt(y)

	return new(big.Float).Quo(k, h)
}
