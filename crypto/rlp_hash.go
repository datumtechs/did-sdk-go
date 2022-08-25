package crypto

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// RlpSHA3Hex returns a hex string with 0x prefix
func RlpSHA3Hex(data interface{}) string {
	h := RlpSHA3(data)
	return h.Hex()
}

func RlpSHA3(data interface{}) (h ethcommon.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, data)
	hw.Sum(h[:0])
	return h
}
