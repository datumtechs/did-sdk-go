package crypto

import "testing"

func Test_LegacyKeccak256SHA3Hex(t *testing.T) {
	t.Log(LegacyKeccak256SHA3Hex("test"))
}

func Test_LegacyKeccak256SHA3(t *testing.T) {
	t.Log(LegacyKeccak256SHA3("test"))
}
