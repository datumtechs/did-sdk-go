package crypto

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"io"
)

// LegacyKeccak256SHA3Hex returns a hex string with 0x prefix
func LegacyKeccak256SHA3Hex(data string) string {
	h := LegacyKeccak256SHA3(data)
	return h.Hex()
}

func LegacyKeccak256SHA3(data string) (h ethcommon.Hash) {
	w := sha3.NewLegacyKeccak256()
	io.WriteString(w, data)
	w.Sum(h[:0])
	return h
}
