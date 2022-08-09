package crypto

import (
	ethhexutl "github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
	"io"
)

// SHA3Hex returns a hex string with 0x prefix
func SHA3Hex(data string) string {
	w := sha3.New256()
	io.WriteString(w, data)
	return ethhexutl.Encode(w.Sum(nil))
}

func SHA3(data string) []byte {
	w := sha3.New256()
	io.WriteString(w, data)
	return w.Sum(nil)
}
