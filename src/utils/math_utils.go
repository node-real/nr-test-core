package utils

import (
	"github.com/shopspring/decimal"
	"math/rand"
)

const (
	bytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytesToLower = "abcdefghijklmnopqrstuvwxyz"
	bytesToUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type MathUtils struct {
}

// 判断一个数是否为偶数
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

func (m *MathUtils) Add(x, y float64) float64 {
	ret, _ := decimal.NewFromFloat(x).Add(decimal.NewFromFloat(y)).Float64()
	return ret
}

func (m *MathUtils) Sub(x, y float64) float64 {
	ret, _ := decimal.NewFromFloat(x).Sub(decimal.NewFromFloat(y)).Float64()
	return ret
}

func (m *MathUtils) Mul(x, y float64) float64 {
	ret, _ := decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y)).Float64()
	return ret
}

func (m *MathUtils) Div(x, y float64) float64 {
	ret, _ := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y)).Float64()
	return ret
}

//func (m *MathUtils) LessThan(x, y float64) bool {
//	return decimal.NewFromFloat(x).LessThan(decimal.NewFromFloat(y))
//}
//
//func (m *MathUtils) GreaterThan(x, y float64) bool {
//	return decimal.NewFromFloat(x).GreaterThan(decimal.NewFromFloat(y))
//}
