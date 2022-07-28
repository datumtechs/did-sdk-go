package vc

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/types/claim"
	"github.com/datumtechs/did-sdk-go/types/proof"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"math/big"
)

/*const (
	TYPE_VC           string = "VerifiableCredential"
	TYPE_VP           string = "VerifiablePresentation"
	PROOF_SALT        string = "salt"
	CLAIM             string = "claim"
	PROOF_DISCLOSURES string = "disclosures"
)*/

type Credential struct {
	Context        string
	Id             string
	PctId          *big.Int
	Type           []string
	Issuer         string // the issuer DID.
	IssuanceDate   string
	ExpirationDate string
	Claim          claim.Claim
	Proof          proof.Proof
	Holder         string // the holder DID.
}

type CredentialWrapper struct {
	Credential *Credential
	// key is the credential field, and value "1" for disclosure to the third party, "0"
	Disclosure map[string]int
}

func (c *Credential) GetCredentialThumbprintWithoutSig(disclosures map[string]int, seed uint64) string {
	rawCredMap := c.GetRawCredentialMap()
	claimHash := c.Claim.GetHash(disclosures, seed)
	rawCredMap["claimHash"] = claimHash
	return ToJson(rawCredMap)
}

func (c *Credential) GetRawCredentialMap() map[string]interface{} {
	cred := make(map[string]interface{})
	cred["context"] = c.Context
	cred["id"] = c.Id
	cred["pctId"] = c.PctId
	cred["issuer"] = c.Issuer
	cred["holder"] = c.Holder
	cred["issuanceDate"] = c.IssuanceDate
	cred["expirationDate"] = c.ExpirationDate
	return cred
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

func (c *Credential) GenerateRawData(claimHash string) string {
	cred := make(map[string]interface{})
	cred["context"] = c.Context
	cred["id"] = c.Id
	cred["pctId"] = c.PctId
	cred["issuer"] = c.Issuer
	cred["holder"] = c.Holder
	cred["issuanceDate"] = c.IssuanceDate
	cred["expirationDate"] = c.ExpirationDate
	cred["claimHash"] = claimHash

	return ToJson(cred)
}

func ToJson(m map[string]interface{}) string {
	dataMap, _ := json.Marshal(m)
	return string(dataMap)
}

type ProofBrief struct {
	CredentialHash ethcommon.Hash
	SignerPubKey   string
	Signature      string
}

func (proof *ProofBrief) AddSignerPubKey(signerPubKey string) {
	if len(proof.SignerPubKey) == 0 {
		proof.SignerPubKey = signerPubKey
	}
	return
}
func (proof *ProofBrief) AddSignature(signature string) {
	if len(proof.Signature) == 0 {
		proof.Signature = signature
	}
	return
}
