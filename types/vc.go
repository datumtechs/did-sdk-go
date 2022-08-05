package types

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/keys/vc"
	log "github.com/sirupsen/logrus"
)

const (
	VERSION = "1.0.0"
)

const (
	CREDENTIAL_TYPE_VC string = "VerifiableCredential"
	CREDENTIAL_TYPE_VP string = "VerifiablePresentation"
)

const (
	DEFAULT_CREDENTIAL_CONTEXT string = "https://www.w3.org/2018/credentials/v1"
)

const (
	VC_EVENT_SIGNERPUBKEY uint8 = 0
	VC_EVENT_SIGNATURE    uint8 = 1
)

type CredentialStatus int8

const (
	Credential_VALID CredentialStatus = iota
	Credential_INVALID
)

func (s CredentialStatus) String() string {
	switch s {
	case Credential_VALID:
		return "Valid"
	case Credential_INVALID:
		return "Invalid"
	default:
		return "NA"
	}
}

type Credential struct {
	Context        string            `json:"context,omitempty"`
	Version        string            `json:"version,omitempty"`
	Id             string            `json:"id,omitempty"`
	Type           []string          `json:"type,omitempty"`
	Issuer         string            `json:"issuer,omitempty"` // the issuer DID.
	IssuanceDate   string            `json:"issuanceDate,omitempty"`
	ExpirationDate string            `json:"expirationDate,omitempty"`
	ClaimData      Claim             `json:"claimData,omitempty"`
	ClaimMeta      map[string]string `json:"claimMeta,omitempty"` //todo: just define as string - pctId
	Proof          Proof             `json:"proof,omitempty"`
	Holder         string            `json:"holder,omitempty"` // the holder DID.
}

type CredentialWrapper struct {
	Credential *Credential
	// key is the credential field, and value "1" for disclosure to the third party, "0"
	Disclosure map[string]int
}

// When seed=0, a random number will be generated as seed.
func (c *Credential) GetDigest(seed uint64) (credentialHash []byte, rootHash string) {
	claimHash, rootHash := c.ClaimData.GetHash(seed)
	credMap := c.ToMap()
	delete(credMap, vckeys.PROOF)
	credMap[vckeys.CLAIM_DATA] = claimHash
	return crypto.SHA3(common.MapToJson(credMap)), rootHash
}

// todo: convert to map by reflect
func (c *Credential) ToMap() map[string]interface{} {
	data, err := json.Marshal(&c)
	if err != nil {
		log.Errorf("cannot Marshal credential, error: %+v", err)
		return nil
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Errorf("cannot Unmarshal credential to Map, error: %+v", err)
		return nil
	}
	return m
}

type EvidenceSignInfo struct {
	CredentialHash string
	SignerPubKey   string
	Signature      string
	Timestamp      string
	Status         string
}

func (proof *EvidenceSignInfo) SupplementSignerPubKey(signerPubKey string) {
	if len(proof.SignerPubKey) == 0 {
		proof.SignerPubKey = signerPubKey
	}
	return
}
func (proof *EvidenceSignInfo) SupplementSignature(signature string) {
	if len(proof.Signature) == 0 {
		proof.Signature = signature
	}
	return
}
