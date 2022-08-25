package types

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/crypto"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

type Presentation struct {
	Context              string        `json:"@context,omitempty"`
	Type                 []string      `json:"type,omitempty"`
	VerifiableCredential []*Credential `json:"verifiableCredential,omitempty"`
	Proof                Proof         `json:"proof,omitempty"`
}

func (p *Presentation) ToRawData() string {
	p.Proof = nil
	b, _ := json.Marshal(p)
	return string(b)
}

func (c *Presentation) GetDigest() ethcommon.Hash {
	return crypto.RlpSHA3(c.ToRawData())
}
