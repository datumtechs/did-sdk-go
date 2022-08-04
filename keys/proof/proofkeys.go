package proofkeys

type ProofKey string

const (
	CREATED            ProofKey = "created"
	TYPE               ProofKey = "type"
	VERIFICATIONMETHOD ProofKey = "verificationMethod" //publicKeyId
	JWS                ProofKey = "jws"
	CHALLENGE          ProofKey = "challenge" //for Presentation
)
