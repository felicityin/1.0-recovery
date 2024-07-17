package common

import (
	"github.com/ipfs/go-log"
)

var Logger = log.Logger("recovery")

func init() {
	if err := log.SetLogLevel("recovery", "info"); err != nil {
		panic(err)
	}
}
