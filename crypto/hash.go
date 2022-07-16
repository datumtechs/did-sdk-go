package crypto

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"io"
)

func SHA3Hex(data string) string {
	w := sha3.New256()
	io.WriteString(w, data)
	return hex.EncodeToString(w.Sum(nil))
}
