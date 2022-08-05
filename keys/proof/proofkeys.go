package proofkeys

type ProofKey string

const (
	CREATED            ProofKey = "created"
	TYPE               ProofKey = "type"
	VERIFICATIONMETHOD ProofKey = "verificationMethod" //publicKeyId
	JWS                ProofKey = "jws"
	SEED               ProofKey = "seed"
	CLAIM_ROOT_HASH    ProofKey = "claimRootHash"
	CHALLENGE          ProofKey = "challenge"   //for Presentation
	DISCLOSURES        ProofKey = "disclosures" //for Presentation
)
