package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type Pct struct {
	Issuer     common.Address
	JsonSchema string
	Extra      []byte
}
