package types

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
	"io"
	"sort"
	"strings"
)

const (
	TYPE_VC           string = "VerifiableCredential"
	TYPE_VP           string = "VerifiablePresentation"
	PROOF_SALT        string = "salt"
	CLAIM             string = "claim"
	PROOF_DISCLOSURES string = "disclosures"

	CLAIM_FIELD_NOT_DISCLOSED int = 0
	CLAIM_FIELD_DISCLOSED     int = 1
	CLAIM_FIELD_EXISTED       int = 2
)

type Credential struct {
	Context        string
	Id             string
	PctId          int
	Issuer         string // the issuer DID.
	IssuanceDate   string
	ExpirationDate string
	Claim          map[string]string
	Proof          map[string]interface{}
	VCType         []string
	Holder         string // the holder DID.
	Version        string
}

type CredentialWrapper struct {
	Credential *Credential
	// key is the credential field, and value "1" for disclosure to the third party, "0"
	Disclosure map[string]int
}

func (c *Credential) GetCredentialThumbprintWithoutSig(disclosures map[string]int) string {
	cc := *c
	cc.Proof = nil
	credMap := cc.ToMap()

	claimHash := getClaimHash(cc.Claim, disclosures)
	credMap[CLAIM] = claimHash
	return ToJson(credMap)
}

func getClaimHash(claim map[string]string, disclosures map[string]int) string {
	if disclosures == nil {
		disclosures = make(map[string]int)
	}

	if len(disclosures) == 0 {
		//每个字段都需要披露
		for key, _ := range claim {
			disclosures[key] = CLAIM_FIELD_DISCLOSED
		}
	}

	for key, _ := range disclosures {
		claim[key] = getSHA3(claim[key])
	}

	keys := make([]string, len(claim))
	i := 0
	for key := range claim {
		keys[i] = key
		i++
	}
	//排序
	keys = sort.StringSlice(keys)

	var sb = new(strings.Builder)
	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteString(claim[key])
	}
	return sb.String()
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

func ToJson(m map[string]interface{}) string {
	dataMap, _ := json.Marshal(m)
	return string(dataMap)
}

func getSHA3(data string) string {
	w := sha3.New256()
	io.WriteString(w, data)
	return hex.EncodeToString(w.Sum(nil))
}

type ProofBrief struct {
	CredentialHash common.Hash
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
