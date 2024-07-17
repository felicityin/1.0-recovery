package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"1.0-recovery/crypto"
	"github.com/mr-tron/base58"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func DecryptSaveData(saveData, role string, address string, pwd string, eccsk []byte) (save []byte, err error) {
	var skBytes []byte

	if role == "shop" {
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("%s-%x", address, eccsk)))
		if pwd != fmt.Sprintf("%x", h.Sum(nil)) {
			return nil, fmt.Errorf("pwd not match with ECCSkHex")
		}
		skBytes = eccsk
	} else {
		pwdBytes, err := base58.Decode(pwd)
		if err != nil {
			return nil, fmt.Errorf("base58 decode pwd err: %s", err.Error())
		}
		skBytes, err = crypto.EciesDecrypt(eccsk, pwdBytes)
		if err != nil {
			return nil, fmt.Errorf("ecies decrypt pwd err: %s", err.Error())
		}
	}

	pkContentBytes, err := base58.Decode(saveData)
	if err != nil {
		return nil, fmt.Errorf("base58 decode keygen data err: %s", err.Error())
	}
	plsBytes, err := crypto.EciesDecrypt(skBytes, pkContentBytes)
	if err != nil {
		return nil, fmt.Errorf("ecies decode pk content err: %s", err.Error())
	}
	return plsBytes, nil
}

func getPrivateKeyfromWIF(wifprivate string) []byte {
	if checkWIF(wifprivate) {
		rawdata := []byte(wifprivate)
		base58decodedata := Base58Decode(rawdata)
		return base58decodedata[1:33]
	}
	return []byte{}

}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		if b == '1' {
			zeroBytes++
		} else {
			break
		}
	}

	payload := input[zeroBytes:]

	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)

		result.Mul(result, big.NewInt(58))

		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()

	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)
	return decoded
}

func checkWIF(wifprivate string) bool {
	rawdata := []byte(wifprivate)
	base58decodedata := Base58Decode(rawdata)

	length := len(base58decodedata)

	if length < 37 {
		Logger.Errorf("wif privkey len < 37")
		return false
	}

	private := base58decodedata[:(length - 4)]

	firstsha := sha256.Sum256(private)

	secondsha := sha256.Sum256(firstsha[:])

	checksum := secondsha[:4]
	orignchecksum := base58decodedata[(length - 4):]
	if bytes.Compare(checksum, orignchecksum) == 0 {
		return true
	}

	return false
}

func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)

	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)

	for _, b := range input {

		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result

}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func CalcWifPrivKey(hexPrivkey string, compressed bool) []byte {
	versionstr := ""
	if compressed {
		versionstr = "80" + hexPrivkey + "01"
	} else {
		versionstr = "80" + hexPrivkey
	}
	privatekey, _ := hex.DecodeString(versionstr)
	firsthash := sha256.Sum256(privatekey)

	secondhash := sha256.Sum256(firsthash[:])

	checksum := secondhash[:4]

	result := append(privatekey, checksum...)

	base58result := Base58Encode(result)
	return base58result
}
