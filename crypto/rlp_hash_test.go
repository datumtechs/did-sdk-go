package crypto

import "testing"

func Test_SHA3Hex(t *testing.T) {
	t.Log(RlpSHA3Hex("test"))
}
