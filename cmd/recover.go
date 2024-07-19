package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	eddsaKeygen "github.com/bnb-chain/tss-lib/eddsa/keygen"
	"github.com/decred/dcrd/dcrec/edwards/v2"

	"recovery-key/common"
	"recovery-key/crypto"
)

type RecoveryInput struct {
	Address string `yaml:"address"`
	Nodes   []Node `yaml:"nodes"`
}

type Node struct {
	Role      string     `yaml:"role"`
	PkContent string     `yaml:"pk_content"`
	Pwd       string     `yaml:"pwd"`
	Eccsk     string     `yaml:"eccsk"`
	Kms       common.Kms `yaml:"kms"`
}

func RecoverKey(paramsPath string) (*big.Int, error) {
	params := loadRecoveryParams(paramsPath)

	checkParams(&params)

	sks := big.NewInt(0)
	for _, node := range params.Nodes {
		sk, n, err := decryptSkSlice(node, params.Address)
		if err != nil {
			return nil, err
		}
		sks.Add(sks, sk)
		sks.Mod(sks, n)
	}

	return sks, nil
}

func loadRecoveryParams(path string) RecoveryInput {
	bytess, err := ioutil.ReadFile(path)
	if err != nil {
		common.Logger.Errorf("load params error: %s", err.Error())
		panic(err)
	}

	var params RecoveryInput
	if err = yaml.UnmarshalStrict(bytess, &params); err != nil {
		common.Logger.Errorf("unmarshal params error: %s", err.Error())
		panic(err)
	}
	return params
}

func checkParams(param *RecoveryInput) error {
	if param.Address == "" {
		return fmt.Errorf("address is empty")
	}
	if len(param.Nodes) == 0 {
		return fmt.Errorf("nodes is empty")
	}
	return nil
}

func decryptSkSlice(node Node, address string) (*big.Int, *big.Int, error) {
	eccsk, err := hex.DecodeString(node.Eccsk)
	if err != nil {
		return nil, nil, fmt.Errorf("hex decode eccsk err: %s", err.Error())
	}

	plsBytes, err := common.DecryptSaveData(node.PkContent, node.Role, address, node.Pwd, eccsk, node.Kms)
	if err != nil {
		return nil, nil, err
	}

	var ecdsaSave *keygen.LocalPartySaveData = &keygen.LocalPartySaveData{}
	var eddsaSave *eddsaKeygen.LocalPartySaveData = &eddsaKeygen.LocalPartySaveData{}

	err = json.Unmarshal(plsBytes, &ecdsaSave)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal pkContent err: %s", err)
	}

	if ecdsaSave.ECDSAPub == nil {
		err := json.Unmarshal(plsBytes, &eddsaSave)
		if err != nil {
			return nil, nil, fmt.Errorf("unmarshal pkContent err: %s", err)
		}
	}

	if ecdsaSave.ECDSAPub != nil {
		return calcEcdsaSkSlice(ecdsaSave), crypto.S256().Params().N, nil
	}
	return calcEddsaSkSlice(eddsaSave), crypto.Edwards().Params().N, nil
}

func calcEcdsaSkSlice(save *keygen.LocalPartySaveData) *big.Int {
	i, err := save.OriginalIndex()
	if err != nil {
		common.Logger.Errorf("invalid index: %d \n", i)
		panic(err)
	}

	wi := save.Xi
	N := crypto.S256().Params().N
	modQ := common.ModInt(N)

	for j := 0; j < 4; j++ {
		if j == i {
			continue
		}
		// big.Int Div is calculated as: a/b = a * modInv(b,q)
		coef := modQ.Mul(save.Ks[j], modQ.ModInverse(new(big.Int).Sub(save.Ks[j], save.Ks[i])))
		wi = modQ.Mul(wi, coef)
	}
	return wi
}

func calcEddsaSkSlice(save *eddsaKeygen.LocalPartySaveData) *big.Int {
	i, err := save.OriginalIndex()
	if err != nil {
		common.Logger.Errorf("invalid index: %d\n", i)
		panic(err)
	}

	wi := save.Xi
	N := edwards.Edwards().N
	modQ := common.ModInt(N)

	for j := 0; j < 4; j++ {
		if j == i {
			continue
		}
		// big.Int Div is calculated as: a/b = a * modInv(b,q)
		coef := modQ.Mul(save.Ks[j], modQ.ModInverse(new(big.Int).Sub(save.Ks[j], save.Ks[i])))
		wi = modQ.Mul(wi, coef)
	}
	return wi
}
