package cmd

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/decred/dcrd/dcrec/edwards/v2"

	"1.0-recovery/common"
	"1.0-recovery/crypto"
)

func MergeKeys(hbcSlice string, shopSlice string, chain string) (*big.Int, string, string, error) {
	hbcPriv, err := hex.DecodeString(hbcSlice)
	if err != nil {
		return nil, "", "", fmt.Errorf("hex decode hbc privkey slice err: %s", err.Error())
	}
	hbcKey := new(big.Int).SetBytes(hbcPriv)

	shopPriv, err := hex.DecodeString(shopSlice)
	if err != nil {
		return nil, "", "", fmt.Errorf("hex decode shop privkey slice err: %s", err.Error())
	}
	shopKey := new(big.Int).SetBytes(shopPriv)

	sk := big.NewInt(0)
	sk.Add(sk, hbcKey)
	sk.Add(sk, shopKey)

	if !common.IsEddsaChain(chain) {
		sk.Mod(sk, crypto.S256().Params().N)
		wifPrivkey := common.CalcWifPrivKey(hex.EncodeToString(sk.Bytes()), false)
		pubkey := crypto.ScalarBaseMult(crypto.S256(), sk)
		pub := ecdsa.PublicKey{
			Curve: pubkey.Curve(),
			X:     pubkey.X(),
			Y:     pubkey.Y(),
		}
		address, err := common.SwitchEcdsaChainAddress(&pub, chain)
		if err != nil {
			return nil, "", "", fmt.Errorf("calc address err: %s", err.Error())
		}
		return sk, string(wifPrivkey), address, nil
	} else {
		sk.Mod(sk, crypto.Edwards().Params().N)
		pubkey := crypto.ScalarBaseMult(crypto.Edwards(), sk)
		pub := edwards.NewPublicKey(pubkey.X(), pubkey.Y())
		address, err := common.SwitchEddsaChainAddress(pub, chain)
		if err != nil {
			return nil, "", "", fmt.Errorf("calc address err: %s", err.Error())
		}
		return sk, "", address, nil
	}
}
