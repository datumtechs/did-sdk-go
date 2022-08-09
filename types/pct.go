package types

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
)

type Pct struct {
	Issuer     ethcommon.Address
	JsonSchema string
	Extra      []byte
}
