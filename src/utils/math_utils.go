package utils

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
	"time"
)

const (
	bytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytesToLower = "abcdefghijklmnopqrstuvwxyz"
	bytesToUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type MathUtils struct {
}

func rdm(n int, data string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = data[rand.Int63()%int64(len(data))]
	}
	return string(b)
}

func (m *MathUtils) RandomString(n int) string {
	return rdm(n, bytes)
}

func (m *MathUtils) RandomStringToLower(n int) string {
	return rdm(n, bytesToLower)
}

func (m *MathUtils) RandomStringToUpper(n int) string {
	return rdm(n, bytesToUpper)
}

func (m *MathUtils) RandHexKey() string {
	key, _ := crypto.GenerateKey()

	keyBytes := crypto.FromECDSA(key)
	hexkey := hexutil.Encode(keyBytes)[2:]
	return hexkey
}
