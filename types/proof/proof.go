package proof

type ProofKey string

const (
	CREATED       ProofKey = "created"
	TYPE          ProofKey = "type"
	PUBLIC_KEY_ID ProofKey = "publicKeyId"
	SIGNATURE     ProofKey = "signature"
)

type Proof map[ProofKey]string
