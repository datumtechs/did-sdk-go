package types

import (
	proofkeys "github.com/datumtechs/did-sdk-go/keys/proof"
	"testing"
)

type Proof2[KEY string, VALUE string | Claim] map[KEY]VALUE

func Test_proof(t *testing.T) {
	var proof2 Proof2[string, string] = map[string]string{string(proofkeys.CREATED): "b"}
	t.Log(proof2)
}
