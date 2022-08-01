package types

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/keys/vc"
	ethcommon "github.com/ethereum/go-ethereum/common"
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

type Credential struct {
	Context        string            `json:"context,omitempty"`
	Version        string            `json:"version,omitempty"`
	Id             string            `json:"id,omitempty"`
	Type           []string          `json:"type,omitempty"`
	Issuer         string            `json:"issuer,omitempty"` // the issuer DID.
	IssuanceDate   string            `json:"issuanceDate,omitempty"`
	ExpirationDate string            `json:"expirationDate,omitempty"`
	ClaimData      Claim             `json:"claimData,omitempty"`
	ClaimMeta      map[string]string `json:"claimMeta,omitempty"`
	Proof          Proof             `json:"proof,omitempty"`
	Holder         string            `json:"holder,omitempty"` // the holder DID.

}

type CredentialWrapper struct {
	Credential *Credential
	// key is the credential field, and value "1" for disclosure to the third party, "0"
	Disclosure map[string]int
}

/*func (c *Credential) GetCredentialThumbprintWithoutSig(disclosures map[string]int, seed uint64) string {
	rawCredMap := c.GetRawCredentialMap()
	claimHash := c.ClaimData.GetHash(disclosures, seed)
	rawCredMap["claimData"] = claimHash
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
*/

func (c *Credential) GetRaw(disclosureMap map[string]int, seed uint64) string {
	claimHash := c.ClaimData.GetHash(disclosureMap, seed)
	credMap := c.ToMap()
	delete(credMap, vckeys.PROOF)
	credMap[vckeys.CLAIM_DATA] = claimHash
	return common.MapToJson(credMap)
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

/*
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
}*/

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
