package common

import (
	"crypto/ecdsa"
	"fmt"

	hcchaincfg "github.com/HcashOrg/hcd/chaincfg"
	"github.com/HcashOrg/hcd/chaincfg/chainec"
	"github.com/HcashOrg/hcd/hcutil"
	btcec "github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

// coin
const (
	BTC      = "btc"
	BTC_TEST = "btc_test"
	LTC      = "ltc"
	DOGE     = "doge"
	ETH      = "eth"
	BCH      = "bch"
	BNB_BSC  = "bnb_bsc"
	HT_HECO  = "ht_heco"
	TRX      = "trx"
	USDT     = "usdt"
	HC       = "hc"
	SOL      = "sol"
)

func IsEddsaChain(chain string) bool {
	return chain == "sol" || chain == "apt" || chain == "dot"
}

func SwitchEcdsaChainAddress(ecdsaPk *ecdsa.PublicKey, chain string) (string, error) {
	var addressStr string
	switch chain {
	case ETH:
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case BNB_BSC:
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case HT_HECO:
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case TRX:
		a := addr.PubkeyToAddress(*ecdsaPk)
		addressStr = a.String()
	case BTC:
		pubkey := btcec.PublicKey(*ecdsaPk)
		pubkeyaddr := &pubkey
		param := &chaincfg.MainNetParams
		pkHash, err := btcutil.NewAddressPubKey(pubkeyaddr.SerializeCompressed(), param)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case BCH:
		pubkey := btcec.PublicKey(*ecdsaPk)
		pubkeyaddr := &pubkey
		pkHash, err := btcutil.NewAddressPubKey(pubkeyaddr.SerializeCompressed(), &BCHParams)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case BTC_TEST:
		pubkey := btcec.PublicKey(*ecdsaPk)
		pubkeyaddr := &pubkey
		pkHash, err := btcutil.NewAddressPubKey(pubkeyaddr.SerializeCompressed(), &chaincfg.TestNet3Params)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case LTC:
		pubkey := btcec.PublicKey(*ecdsaPk)
		pubkeyaddr := &pubkey
		pkHash, err := btcutil.NewAddressPubKey(pubkeyaddr.SerializeCompressed(), &LTCParams)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case DOGE:
		pubkey := btcec.PublicKey(*ecdsaPk)
		pubkeyaddr := &pubkey
		pkHash, err := btcutil.NewAddressPubKey(pubkeyaddr.SerializeCompressed(), &DOGEParams)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case USDT:
		pubkey := btcec.PublicKey(*ecdsaPk)
		pubkeyaddr := &pubkey
		param := &chaincfg.MainNetParams
		pkHash, err := btcutil.NewAddressPubKey(pubkeyaddr.SerializeCompressed(), param)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case HC:
		pubKey := ecrypto.CompressPubkey(ecdsaPk)
		pubKeyHash := hcutil.Hash160(pubKey)
		param := &hcchaincfg.MainNetParams
		addr, err := hcutil.NewAddressPubKeyHash(pubKeyHash,
			param, chainec.ECTypeSecp256k1)
		if err != nil {
			return "", err
		}
		addressStr = addr.EncodeAddress()
	default:
		return "", fmt.Errorf("ecdsa, unsupport chain type for %s", chain)
	}
	return addressStr, nil
}

func SwitchEddsaChainAddress(publicKey *edwards.PublicKey, chain string) (addressStr string, err error) {
	switch chain {
	case SOL:
		addressStr = base58.Encode(publicKey.Serialize())
	default:
		return "", fmt.Errorf("eddsa, unsupported chain: %s", chain)
	}
	return addressStr, nil
}
