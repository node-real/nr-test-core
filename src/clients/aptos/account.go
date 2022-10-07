package aptos

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

type Account struct {
	PrivateKey ed25519.PrivateKey
	Publickey  ed25519.PublicKey
	address    string
}

func InitAccount() *Account {
	publickey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	instance := new(Account)
	instance.PrivateKey = privateKey
	instance.Publickey = publickey
	return instance
}

func InitAccountWithSeed(seed string) *Account {
	seeds, _ := hex.DecodeString(seed)
	publickey, privateKey, err := ed25519.GenerateKey(bytes.NewReader(seeds))
	if err != nil {
		panic(err)
	}
	instance := new(Account)
	instance.PrivateKey = privateKey
	instance.Publickey = publickey
	return instance
}

func (a *Account) Address() string {
	return a.Auth_key()
}

func (a *Account) Auth_key() string {
	//auth_key = auth_key = sha3-256(pubkey_A | 0x00)
	if a.address == "" {
		hash := sha3.New256()
		hash.Write(append(a.Publickey, []byte{0x00}...))
		a.address = hex.EncodeToString(hash.Sum(nil))
	}
	return a.address
}

func (a *Account) Pub_key() string {
	//"""Returns the public key for the associated account"""
	return hex.EncodeToString(a.Publickey)
}

func (a *Account) GetSeed() string {
	return hex.EncodeToString(a.PrivateKey.Seed())

}
