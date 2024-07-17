package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/bnb-chain/tss-lib/common"

	"1.0-recovery/cmd"
)

func main() {
	recoverCmd := flag.NewFlagSet("recover", flag.ExitOnError)
	inputPath := recoverCmd.String("i", "./test/shop_input.yaml", "The path of input parmas")

	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	chainNam := mergeCmd.String("chain", "eth", "chain")
	hbcSlice := mergeCmd.String("hbc", "", "HBC private key slice")
	shopSlice := mergeCmd.String("shop", "", "shop private key slice")

	balanceCmd := flag.NewFlagSet("balance", flag.ExitOnError)
	address := balanceCmd.String("addr", "", "address")
	coin := balanceCmd.String("coin", "", "Coin contract address. For sol, refer to https://solscan.io/leaderboard/token")
	chain := balanceCmd.String("chain", "sol", "chain")
	url := balanceCmd.String("url", "https://api.mainnet-beta.solana.com", "url")

	transferCmd := flag.NewFlagSet("transfer", flag.ExitOnError)
	fromkey := transferCmd.String("fromkey", "", "private key")
	toAddress := transferCmd.String("to", "", "address")
	amount := transferCmd.String("amount", "", "amount")
	coinAddress := transferCmd.String("coin", "", "Coin contract address. For sol, refer to https://solscan.io/leaderboard/token")
	chainName := transferCmd.String("chain", "sol", "chain")
	chainUrl := transferCmd.String("url", "https://api.mainnet-beta.solana.com", "url")

	if len(os.Args) < 2 {
		fmt.Println("expected 'recover', 'balance' or 'transfer' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "recover":
		recoverCmd.Parse(os.Args[2:])
		sk, err := cmd.RecoverKey(*inputPath)
		if err != nil {
			common.Logger.Errorf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("sk: %s\n ", hex.EncodeToString(sk.Bytes()))
	case "merge":
		mergeCmd.Parse(os.Args[2:])
		sk, addr, err := cmd.MergeKeys(*hbcSlice, *shopSlice, *chainNam)
		if err != nil {
			common.Logger.Errorf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("privkey: %s\n", sk)
		fmt.Printf("address: %s\n", addr)
	case "balance":
		balanceCmd.Parse(os.Args[2:])
		amount, err := cmd.GetBalance(*chain, *url, *address, *coin)
		if err != nil {
			common.Logger.Errorf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("balance: %s\n", amount.Balance)
		fmt.Printf("decimals: %s\n", amount.Decimals)
		fmt.Printf("amount: %s\n", amount.Amount)
	case "transfer":
		transferCmd.Parse(os.Args[2:])

		txHash, err := cmd.Transfer(*chainName, *chainUrl, *fromkey, *toAddress, *amount, *coinAddress)
		if err != nil {
			common.Logger.Errorf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("tx: %s/%s\n", cmd.Scan(*chainName), txHash)
	default:
		fmt.Println("expected 'recover', 'balance' or 'transfer' subcommands")
		os.Exit(1)
	}
}
