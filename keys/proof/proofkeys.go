package proofkeys

type ProofKey string

const (
	CREATED            ProofKey = "created"
	TYPE               ProofKey = "type"
	VERIFICATIONMETHOD ProofKey = "verificationMethod" //publicKeyId
	JWS                ProofKey = "jws"
	SIGNATURE          ProofKey = "signature"
	SEED               ProofKey = "seed"
	ROOT_HASH          ProofKey = "rootHash"
)
