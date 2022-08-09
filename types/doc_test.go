package types

import (
	platoncommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_regexp(t *testing.T) {
	s := "did:pid:1334ada"
	address := GetAddressFromDid(s)

	a := assert.New(t)
	a.Equal("1334ada", address)
}
func Test_ParseToAddress(t *testing.T) {
	did := "did:pid:lat1cq9svdd8vc83u74relncn6cyxywr5mjqccqlea"
	addr, err := ParseToAddress(did)
	a := assert.New(t)
	if a.NoError(err) {
		a.Equal("lat1cq9svdd8vc83u74relncn6cyxywr5mjqccqlea", platoncommon.HexToAddress(addr.Hex()).String())
	}
}
