package aptos

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/node-real/nr-test-core/src/invokers/http"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"time"
)

type TransactionBody struct {
	Sender                  string     `json:"sender"`
	SequenceNumber          string     `json:"sequence_number"`
	MaxGasAmount            string     `json:"max_gas_amount"`
	GasUnitPrice            string     `json:"gas_unit_price"`
	GasCurrencyCode         string     `json:"gas_currency_code"`
	ExpirationTimestampSecs string     `json:"expiration_timestamp_secs"`
	Payload                 *Payload   `json:"payload"`
	Signature               *Signature `json:"signature"`
}

type TransactionSignBody struct {
	Sender         string `json:"sender"`
	SequenceNumber string `json:"sequence_number"`
	MaxGasAmount   string `json:"max_gas_amount"`
	GasUnitPrice   string `json:"gas_unit_price"`
	//GasCurrencyCode         string   `json:"gas_currency_code"`
	ExpirationTimestampSecs string   `json:"expiration_timestamp_secs"`
	Payload                 *Payload `json:"payload"`
}

type Signature struct {
	Signtype  string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type Payload struct {
	// Function Format: {address}::{module name}::{function name}
	//
	//Both module name and function name are case-sensitive.
	Paytype        string    `json:"type"`
	Function       string    `json:"function,omitempty"`
	Type_arguments []string  `json:"type_arguments,omitempty"`
	Arguments      []string  `json:"arguments,omitempty"`
	Modules        []*Module `json:"modules,omitempty"`
}

//type Function struct {
//	Modulef *Modulef `json:"module"`
//	Name    string   `json:"name"`
//}

type Modulef struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
type Module struct {
	// Function Format: {address}::{module name}::{function name}
	//
	//Both module name and function name are case-sensitive.
	Bytecode string `json:"bytecode"`
}

func buildTransignBody(sender *Account, payload *Payload, sequenceNumber string) *TransactionSignBody {

	tsb := new(TransactionSignBody)

	tsb.Sender = "0x" + sender.Address()
	tsb.SequenceNumber = sequenceNumber
	tsb.MaxGasAmount = "2000"
	tsb.GasUnitPrice = "1"
	//tsb.GasCurrencyCode = "XUS"
	tsb.ExpirationTimestampSecs = strconv.FormatInt(time.Now().Unix()+6000, 10)
	tsb.Payload = payload

	return tsb
}

func buildTranBodyFromSign(sender *Account, signbody *TransactionSignBody, signres *http.Response) *TransactionBody {
	tsb := new(TransactionBody)
	tsb.Signature = buildSign(sender, signres)
	tsb.Sender = signbody.Sender
	tsb.SequenceNumber = signbody.SequenceNumber
	tsb.MaxGasAmount = signbody.MaxGasAmount
	tsb.GasUnitPrice = signbody.GasUnitPrice
	//tsb.GasCurrencyCode = signbody.GasCurrencyCode
	tsb.ExpirationTimestampSecs = signbody.ExpirationTimestampSecs
	tsb.Payload = signbody.Payload
	return tsb

}

func buildSign(account *Account, signres *http.Response) *Signature {
	sign := new(Signature)

	message := http.GetBodyParam(signres, "$")
	bytes, _ := hex.DecodeString(message[2:])
	signs, _ := account.PrivateKey.Sign(rand.Reader, bytes, crypto.Hash(0))
	sign.Signature = "0x" + hex.EncodeToString(signs)
	sign.Signtype = "ed25519_signature"
	sign.PublicKey = "0x" + account.Pub_key()

	return sign
}

// todo add more payload
func buildPayload(receiver string, amount string) *Payload {
	payload := new(Payload)
	payload.Paytype = "script_function_payload"
	//payload.Function = "{\"module\":{\"address\": \"0x6353a19e0826be9a9b1071e56feb82256e657660bde2eca0fadeb8815ea2f262\", \"name\": \"coin\"},\"name\": \"transfer\"}"
	payload.Function = "0x1::coin::transfer"
	//payload.Function = new(Function)
	//payload.Function.Name = "transfer"
	//payload.Function.Modulef = new(Modulef)
	//payload.Function.Modulef.Name = "coin"
	//payload.Function.Modulef.Address = "0x1"
	payload.Type_arguments = []string{"0x1::aptos_coin::AptosCoin"}
	payload.Arguments = append(payload.Arguments, "0x"+receiver, amount)
	return payload

}

func buildCreateAccountPayload(address string) *Payload {
	payload := new(Payload)
	payload.Paytype = "script_function_payload"
	payload.Function = "0x1::create_account::Account"
	//payload.Function = new(Function)
	//payload.Function.Name = "create_account"
	//payload.Function.Modulef = new(Modulef)
	//payload.Function.Modulef.Name = "Account"
	//payload.Function.Modulef.Address = address
	payload.Type_arguments = []string{}
	payload.Arguments = append(payload.Arguments, address)
	return payload

}

func buildModulePayload(path string) *Payload {
	payload := new(Payload)
	payload.Paytype = "module_bundle_payload"
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	payload.Modules = make([]*Module, 1)
	payload.Modules[0] = new(Module)
	payload.Modules[0].Bytecode = "0x" + hex.EncodeToString(content)
	return payload

}

func (t *TransactionBody) getJsonBody() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	Body := string(bytes)
	publicKey := gjson.Get(Body, "signature.public_key").String()
	signature := gjson.Get(Body, "signature.signature").String()
	p, _ := new(big.Int).SetString(publicKey, 16)
	s, _ := new(big.Int).SetString(signature, 16)
	Body, err = sjson.Set(Body, "signature.public_key", p)
	Body, err = sjson.Set(Body, "signature.signature", s)
	return Body, nil
}
