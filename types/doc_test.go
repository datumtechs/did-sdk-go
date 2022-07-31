package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_regexp(t *testing.T) {
	s := "did:pid:1334ada"
	address := GetAddressFromDid(s)

	a := assert.New(t)
	a.Equal("1334ada", address)
}
