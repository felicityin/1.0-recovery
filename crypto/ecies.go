package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func EciesDecrypt(sk []byte, data []byte) ([]byte, error) {
	d := new(big.Int).SetBytes(sk)
	pk := &ecdsa.PrivateKey{
		D: d,
	}
	pk.Curve = crypto.S256()
	pk.X, pk.Y = pk.Curve.ScalarBaseMult(d.Bytes())

	prk2 := ecies.ImportECDSA(pk)

	dataBytes, err := ECCDecrypt(data, *prk2)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

// base58编码
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func ECCEncrypt(pt []byte, puk ecies.PublicKey) ([]byte, error) {
	ct, err := ecies.Encrypt(rand.Reader, &puk, pt, nil, nil)
	return ct, err
}

func ECCDecrypt(ct []byte, prk ecies.PrivateKey) ([]byte, error) {
	pt, err := prk.Decrypt(ct, nil, nil)
	return pt, err
}
