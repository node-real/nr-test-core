package utils

import (
	"bufio"
	"github.com/node-real/nr-test-core/src/log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Utils struct {
}

func (utils *Utils) GetStringInBetween(str string, start string, end string) (result string) {
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

func (utils *Utils) ConvertStrToInt(intStr string) (int, error) {
	i, err := strconv.Atoi(intStr)
	return i, err
}

func (utils *Utils) ReadFileToLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func (utils *Utils) WriteFileWithLines(dataItems []string, filePath string) {
	var file *os.File
	var err1 error
	if utils.CheckFileIsExist(filePath) {
		file, err1 = os.OpenFile(filePath, os.O_APPEND, 0666)
	} else {
		file, err1 = os.Create(filePath)
	}
	for _, item := range dataItems {
		file.WriteString(item)
		file.WriteString("\n")
	}
	if err1 != nil {
		log.Fatal(err1)
	}
	defer file.Close()
	file.Sync()
}

func (utils *Utils) CheckFileIsExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (utils *Utils) HexToBigInt(hex string) *big.Int {
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

func (utils *Utils) BigIntToHexString(bigInt *big.Int) string {
	if nil == bigInt {
		return "0x0"
	}
	hex_value := strconv.FormatInt(bigInt.Int64(), 16)
	return "0x" + hex_value
}

func (utils *Utils) Int64ToHexString(intValue int64) string {
	hex_value := strconv.FormatInt(intValue, 16)
	return "0x" + hex_value
}

func (utils *Utils) StringToBigInt(string2 string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(string2, 10)
	if nil == n {
		n = new(big.Int).SetUint64(0)
	}
	return n
}

func (utils *Utils) StringToHexString(string2 string) string {
	big1 := utils.StringToBigInt(string2)
	return utils.BigIntToHexString(big1)
}
