package types

import (
	"github.com/datumtechs/did-sdk-go/keys/proof"
	"strconv"
)

type Proof map[proofkeys.ProofKey]interface{}

func (p Proof) GetSeed() (uint64, error) {
	seedString := p[proofkeys.SEED].(string)
	return strconv.ParseUint(seedString, 10, 64)
}
