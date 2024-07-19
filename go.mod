module recovery-key

go 1.21.4

replace github.com/btcsuite/btcutil/hdkeychain v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcuti/hdkeychain v1.1.3

replace github.com/btcsuite/btcd/btcec => ./package/github.com/btcsuite/btcd/btcec/v1

replace github.com/btcsuite/btcd/btcec/v2 => ./package/github.com/btcsuite/btcd/btcec/v2

require (
	github.com/HcashOrg/bliss v0.0.0-20180719035130-f5d53c2a9b7d // indirect
	github.com/btcsuite/btcd/btcec v0.0.0-00010101000000-000000000000
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/dchest/blake256 v1.1.0 // indirect
	github.com/shengdoushi/base58 v1.0.0 // indirect
)

require (
	github.com/bnb-chain/tss-lib v1.5.0
	github.com/decred/dcrd/dcrec/edwards/v2 v2.0.3
	github.com/mr-tron/base58 v1.2.0
)

require (
	filippo.io/edwards25519 v1.1.0
	github.com/alecthomas/gometalinter v3.0.0+incompatible
	github.com/blocto/solana-go-sdk v1.30.0
	github.com/bnb-chain/edwards25519 v0.0.0-20231030070956-6796d47b70ba
	github.com/btcsuite/btcd v0.23.0
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/fbsobreira/gotron-sdk v0.0.0-20230907131216-1e824406fe8c
	github.com/near/borsh-go v0.3.2-0.20220516180422-1ff87d108454
	github.com/portto/aptos-go-sdk v0.0.0-20230807103729-9a5201cad72f
	github.com/shopspring/decimal v1.4.0
	github.com/stretchr/testify v1.9.0
	github.com/the729/lcs v0.1.5
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hasura/go-graphql-client v0.9.1 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)

require (
	github.com/HcashOrg/hcd v0.0.0-20180816055255-f68c5e6e35cb
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.790
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/go-ethereum v1.13.5
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/holiman/uint256 v1.3.0 // indirect
	github.com/ipfs/go-log v1.0.5
	github.com/ipfs/go-log/v2 v2.1.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/otiai10/primes v0.0.0-20180210170552-f6d2a1ba97c4 // indirect
	github.com/pkg/errors v0.9.1
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)
