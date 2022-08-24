package types

import (
	"bytes"
	"fmt"
	platoncommon "github.com/PlatONnetwork/PlatON-Go/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	DID_PARTERN_PREFIX = "did:pid"
)

type PublicKeyType string

const (
	PublicKey_RSA       PublicKeyType = "RSA"
	PublicKey_SECP256K1 PublicKeyType = "Secp256k1"
)

type PublicKeyStatus int

const (
	PublicKey_VALID PublicKeyStatus = iota
	PublicKey_INVALID
)

func (s PublicKeyStatus) String() string {
	switch s {
	case PublicKey_VALID:
		return "0"
	case PublicKey_INVALID:
		return "1"
	default:
		return "NA"
	}
}

type DocumentStatus int8

const (
	DOC_ACTIVATION DocumentStatus = iota
	DOC_DEACTIVATION
)

func (s DocumentStatus) String() string {
	switch s {
	case DOC_ACTIVATION:
		return "activation"
	case DOC_DEACTIVATION:
		return "deactivation"
	default:
		return "NA"
	}
}

const (
	DOC_EVENT_CREATE   uint8 = 0
	DOC_EVEN_PUBLICKEY uint8 = 1
	DOC_EVEN_SERVICE   uint8 = 1
)

const (
	SEPARATOR_PIPELINE     = "|"
	SEPARATOR_PUBLICKEY_ID = "#keys-"
)

type DidPublicKey struct {
	//did#keys-idx
	Id string
	//公钥类型
	Type string
	//公钥16进制字符串
	PublicKey string //0x prefix
	//公钥是否撤消 0: valid; 1: invalid
	Status PublicKeyStatus
}

type DidService struct {
	Id              string
	Type            string
	ServiceEndpoint string
	Revoked         bool
}

type DidDocument struct {
	Context   string `json:"@context"`
	Id        string
	Version   string
	PublicKey []*DidPublicKey //0x prefix
	Service   []*DidService
	Proof     Proof //0x prefix
	Created   string
	Updated   string
	Status    DocumentStatus
}

// SupplementDidPublicKey supplement a DidPublicKey to current DidDocument if the public key does not exist.
func (doc *DidDocument) SupplementDidPublicKey(pubKey *DidPublicKey) {
	if doc.PublicKey == nil {
		doc.PublicKey = []*DidPublicKey{pubKey}
		return
	}
	if !doc.IsPublicKeyIdOrPublicKeyExist(pubKey.Id, pubKey.PublicKey) {
		doc.PublicKey = append(doc.PublicKey, pubKey)
	}
	return
}

func (doc *DidDocument) IsPublicKeyIdOrPublicKeyExist(publicKeyId string, publicKey string) (exist bool) {
	/*if doc.PublicKey == nil || len(doc.PublicKey) == 0 {
		return false
	}*/
	for _, item := range doc.PublicKey {
		//if item.Id == publicKeyId || item.PublicKey == publicKey {
		if item.Id == publicKeyId || bytes.Equal(ethcommon.FromHex(item.PublicKey), ethcommon.FromHex(publicKey)) {
			return true
		}
	}
	return false
}

func (doc *DidDocument) SupplementService(services []*DidService) {
	if doc.Service == nil {
		doc.Service = services
		return
	}
	var existing bool
	for _, prev := range services {
		existing = false
		for _, element := range doc.Service {
			if element.Id != prev.Id {
				existing = true
			}
		}
		if !existing {
			doc.Service = append(doc.Service, prev)
		}
	}
	return
}

func (doc *DidDocument) FindDidPublicKeyByDidPublicKeyId(didPublicKeyId string) *DidPublicKey {
	if doc.PublicKey == nil || len(doc.PublicKey) == 0 {
		return nil
	}
	for _, didPubKey := range doc.PublicKey {
		if didPubKey.Id == didPublicKeyId {
			return didPubKey
		}
	}
	return nil
}

func (doc *DidDocument) FindDidPublicKeyByPublicKey(publicKeyHex string) *DidPublicKey {
	if doc.PublicKey == nil || len(doc.PublicKey) == 0 {
		return nil
	}
	for _, didPubKey := range doc.PublicKey {
		//if didPubKey.PublicKey == publicKeyHex {
		if bytes.Equal(ethcommon.FromHex(didPubKey.PublicKey), ethcommon.FromHex(publicKeyHex)) {
			return didPubKey
		}
	}
	return nil
}
func BuildDid(address ethcommon.Address) string {
	return fmt.Sprintf("did:pid:%s", platoncommon.Address(address).Bech32())
}

func ExtractAddress(did string) string {
	return string(([]byte(did))[len("did:pid:")-1])
}

func ParseToAddress(did string) (ethcommon.Address, error) {

	addr := string(([]byte(did))[len("did:pid:"):])
	if platoncommon.IsBech32Address(addr) {
		platonAddress, err := platoncommon.Bech32ToAddress(addr)
		if err != nil {
			log.WithError(err).Errorf("failed to parse did: %s", did)
			return ethcommon.Address{}, err
		} else {
			return ethcommon.Address(platonAddress), nil
		}
	} else {
		return ethcommon.HexToAddress(addr), nil
	}
}

func GetAddressFromDid(did string) string {
	//regExp := regexp.MustCompile("(?<=\bdid:pid:)[a-zA-Z0-9]*")
	ids := strings.Split(did, ":")
	if len(ids) == 3 {
		return ids[2]
	}
	return ""

}

func BuildPublicKeyId(did string, idx string) string {
	builder := strings.Builder{}
	builder.WriteString(did)
	builder.WriteString(SEPARATOR_PUBLICKEY_ID)
	builder.WriteString(idx)
	return builder.String()
}

func BuildFieldValueOfPublicKey(pubKey string, pubKeyType PublicKeyType, index string, status PublicKeyStatus) string {
	builder := strings.Builder{}
	builder.WriteString(pubKey)
	builder.WriteString(SEPARATOR_PIPELINE)
	builder.WriteString(string(pubKeyType))
	builder.WriteString(SEPARATOR_PIPELINE)
	builder.WriteString(index)
	builder.WriteString(SEPARATOR_PIPELINE)
	builder.WriteString(status.String())
	return builder.String()
}

func EventToDidPublicKey(did string, eventValue string) *DidPublicKey {
	items := strings.Split(eventValue, SEPARATOR_PIPELINE)

	didPublicKey := new(DidPublicKey)
	didPublicKey.Id = BuildPublicKeyId(did, items[2])
	didPublicKey.PublicKey = items[0]
	didPublicKey.Type = items[1]
	status, err := strconv.Atoi(items[3])
	if err != nil {
		log.WithError(err).Errorf("failed to parse public key status in event, status:%s", items[3])
		didPublicKey.Status = PublicKey_INVALID
	} else {
		didPublicKey.Status = PublicKeyStatus(status)
	}

	return didPublicKey
}
