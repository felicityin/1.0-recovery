package common

import (
	"crypto/ecdsa"
	"fmt"

	hcchaincfg "github.com/HcashOrg/hcd/chaincfg"
	"github.com/HcashOrg/hcd/chaincfg/chainec"
	"github.com/HcashOrg/hcd/hcutil"
	btcec "github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

// chain name
const (
	BitcoinChain     = "Bitcoin"
	EthereumChain    = "Ethereum"
	TronChain        = "Tron"
	BSCChain         = "BSC"
	BitcoinCashChain = "Bitcoin Cash"
	DogeChain        = "Doge"
	LitecoinChain    = "Litecoin"
	HecoChain        = "Heco"
	PolygonChain     = "Polygon"
	ArbitrumChain    = "Arbitrum"
	PolkadotChain    = "Polkadot"
	AptostChain      = "Aptos"
	SolanaChain      = "Solana"
	BaseChain        = "Base Chain"
)

// coin
const (
	BTC           = "btc"
	BTC_TEST      = "btc_test"
	LTC           = "ltc"
	DOGE          = "doge"
	ETH           = "eth"
	BCH           = "bch"
	BNB_BSC       = "bnb_bsc"
	HT_HECO       = "ht_heco"
	TRX           = "trx"
	MATIC_POLYGON = "matic_polygon"
	ETH_ARBITRUM  = "eth_arbitrum"
	APT           = "apt"
	SOL           = "sol"
	DOT           = "dot"
	USDT          = "usdt"
	HC            = "hc"
	DASH          = "dash"
	DCR           = "dcr"
	RVN           = "rvn"
	OKT           = "okt"
	CMP           = "cmp"
	FTM           = "ftm"
	SMARTBCH      = "smartbch"
	ETH_AURORA    = "eth_aurora"
	WEMIX         = "wemix"
	GDCC          = "gdcc"
	ETH_ZKSYNC    = "eth_zksync"
	ETHG          = "ethg"
	CORE          = "core"
	MBE           = "mbe"
	ETHW          = "ethw"
	REI           = "rei"
	ETH_OPTIMISM  = "eth_optimism"
	MOVR          = "movr"
	AVAX_C        = "avax_c"
	ETH_BASE      = "eth_base"
)

func IsEddsaChain(chain string) bool {
	return chain == "sol" || chain == "apt" || chain == "dot"
}

func SwitchEcdsaChainAddress(ecdsaPk *ecdsa.PublicKey, chain string) (string, error) {
	var addressStr string
	switch chain {
	case "eth":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "bnb_bsc":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "ht_heco":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "trx":
		a := addr.PubkeyToAddress(*ecdsaPk)
		addressStr = a.String()
	case "btc":
		var xFieldVal btcec.FieldVal
		var yFieldVal btcec.FieldVal
		if overflow := xFieldVal.SetByteSlice(ecdsaPk.X.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.X.Bytes()) overflow: %x", ecdsaPk.X.Bytes())
			panic(err)
		}
		if overflow := yFieldVal.SetByteSlice(ecdsaPk.Y.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.Y.Bytes()) overflow: %x", ecdsaPk.Y.Bytes())
			panic(err)
		}
		btcecPubkey := btcec.NewPublicKey(&xFieldVal, &yFieldVal)
		param := &chaincfg.MainNetParams
		pkHash, err := btcutil.NewAddressPubKey(btcecPubkey.SerializeCompressed(), param)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case "btc_test":
		var xFieldVal btcec.FieldVal
		var yFieldVal btcec.FieldVal
		if overflow := xFieldVal.SetByteSlice(ecdsaPk.X.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.X.Bytes()) overflow: %x", ecdsaPk.X.Bytes())
			panic(err)
		}
		if overflow := yFieldVal.SetByteSlice(ecdsaPk.Y.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.Y.Bytes()) overflow: %x", ecdsaPk.Y.Bytes())
			panic(err)
		}
		btcecPubkey := btcec.NewPublicKey(&xFieldVal, &yFieldVal)
		pkHash, err := btcutil.NewAddressPubKey(btcecPubkey.SerializeCompressed(), &chaincfg.TestNet3Params)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case "ltc":
		var xFieldVal btcec.FieldVal
		var yFieldVal btcec.FieldVal
		if overflow := xFieldVal.SetByteSlice(ecdsaPk.X.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.X.Bytes()) overflow: %x", ecdsaPk.X.Bytes())
			panic(err)
		}
		if overflow := yFieldVal.SetByteSlice(ecdsaPk.Y.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.Y.Bytes()) overflow: %x", ecdsaPk.Y.Bytes())
			panic(err)
		}
		btcecPubkey := btcec.NewPublicKey(&xFieldVal, &yFieldVal)
		pkHash, err := btcutil.NewAddressPubKey(btcecPubkey.SerializeCompressed(), &LTCParams)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case "doge":
		var xFieldVal btcec.FieldVal
		var yFieldVal btcec.FieldVal
		if overflow := xFieldVal.SetByteSlice(ecdsaPk.X.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.X.Bytes()) overflow: %x", ecdsaPk.X.Bytes())
			panic(err)
		}
		if overflow := yFieldVal.SetByteSlice(ecdsaPk.Y.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.Y.Bytes()) overflow: %x", ecdsaPk.Y.Bytes())
			panic(err)
		}
		btcecPubkey := btcec.NewPublicKey(&xFieldVal, &yFieldVal)
		pkHash, err := btcutil.NewAddressPubKey(btcecPubkey.SerializeCompressed(), &DOGEParams)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case "usdt":
		var xFieldVal btcec.FieldVal
		var yFieldVal btcec.FieldVal
		if overflow := xFieldVal.SetByteSlice(ecdsaPk.X.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.X.Bytes()) overflow: %x", ecdsaPk.X.Bytes())
			panic(err)
		}
		if overflow := yFieldVal.SetByteSlice(ecdsaPk.Y.Bytes()); overflow {
			err := fmt.Errorf("xFieldVal.SetByteSlice(pk.Y.Bytes()) overflow: %x", ecdsaPk.Y.Bytes())
			panic(err)
		}
		btcecPubkey := btcec.NewPublicKey(&xFieldVal, &yFieldVal)
		param := &chaincfg.MainNetParams
		pkHash, err := btcutil.NewAddressPubKey(btcecPubkey.SerializeCompressed(), param)
		if err != nil {
			return "", err
		}
		addressStr = pkHash.EncodeAddress()
	case "hc":
		pubKey := ecrypto.CompressPubkey(ecdsaPk)
		pubKeyHash := hcutil.Hash160(pubKey)
		param := &hcchaincfg.MainNetParams
		addr, err := hcutil.NewAddressPubKeyHash(pubKeyHash,
			param, chainec.ECTypeSecp256k1)
		if err != nil {
			return "", err
		}
		addressStr = addr.EncodeAddress()
	case "bch":
		var err error
		addressStr, err = makeBtcAddress(ecdsaPk, &BCHParams)
		if err != nil {
			return "", err
		}
	case "dash":
		var err error
		addressStr, err = makeBtcAddress(ecdsaPk, &DASHParams)
		if err != nil {
			return "", err
		}
	case "dcr":
		var err error
		addressStr, err = makeBtcAddress(ecdsaPk, &DCRParams)
		if err != nil {
			return "", err
		}
	case "rvn":
		var err error
		addressStr, err = makeBtcAddress(ecdsaPk, &RVNParams)
		if err != nil {
			return "", err
		}
	case "okt":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "cmp":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "ftm":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "smartbch":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "eth_aurora":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "wemix":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "gdcc":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "eth_zksync":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "ethg":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "core":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "mbe":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "ethw":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "rei":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "eth_arbitrum":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "eth_optimism":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "movr":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "avax_c":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	case "matic_polygon":
		address := ecrypto.PubkeyToAddress(*ecdsaPk)
		addressStr = address.Hex()
	default:
		return "", fmt.Errorf("ecdsa, unsupported chain%s", chain)
	}
	return addressStr, nil
}

func SwitchEddsaChainAddress(publicKey *edwards.PublicKey, chain string) (addressStr string, err error) {
	switch chain {
	case "sol":
		addressStr = base58.Encode(publicKey.Serialize())
	default:
		return "", fmt.Errorf("eddsa, unsupported chain: %s", chain)
	}
	return addressStr, nil
}

func makeBtcAddress(ecdsaPk *ecdsa.PublicKey, params *chaincfg.Params) (addressStr string, err error) {
	var xFieldVal btcec.FieldVal
	var yFieldVal btcec.FieldVal
	if overflow := xFieldVal.SetByteSlice(ecdsaPk.X.Bytes()); overflow {
		err := fmt.Errorf("xFieldVal.SetByteSlice(pk.X.Bytes()) overflow: %x", ecdsaPk.X.Bytes())
		panic(err)
	}
	if overflow := yFieldVal.SetByteSlice(ecdsaPk.Y.Bytes()); overflow {
		err := fmt.Errorf("xFieldVal.SetByteSlice(pk.Y.Bytes()) overflow: %x", ecdsaPk.Y.Bytes())
		panic(err)
	}
	btcecPubkey := btcec.NewPublicKey(&xFieldVal, &yFieldVal)
	pkHash, err := btcutil.NewAddressPubKey(btcecPubkey.SerializeCompressed(), params)
	if err != nil {
		return "", err
	}
	addressStr = pkHash.EncodeAddress()
	return addressStr, nil
}
