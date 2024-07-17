package address

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var NetWorkByteMap = map[string]byte{
	"DOT": 0x00,
	"KSM": 0x00,
}

func AddressDecode(addr string, opts ...interface{}) ([]byte, error) {
	data, err := addressEncoder.Base58Decode(addr, addressEncoder.NewBase58Alphabet(addressEncoder.BTCAlphabet))
	if err != nil {
		return nil, err
	}
	pubkey := data[1 : len(data)-2]
	return pubkey, nil
}
