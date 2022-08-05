package types

import proofkeys "github.com/datumtechs/did-sdk-go/keys/proof"

type Proof map[proofkeys.ProofKey]interface{}
