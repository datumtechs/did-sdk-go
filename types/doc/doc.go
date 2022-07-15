package doc

import (
	"encoding/json"
	"fmt"
	"github.com/datumtechs/did-sdk-go/types/proof"
	"strings"
)

type PublicKeyType string

const (
	PublicKey_RSA       PublicKeyType = "RSA"
	PublicKey_SECP256K1 PublicKeyType = "Secp256k1"
)

type PublicKeyStatus string

type DidPublicKey struct {
	//公钥ID
	Id string
	//公钥类型
	Type string

	//公钥16进制字符串
	PublicKey string
	//公钥是否撤消
	Revoked bool
}

type DidService struct {
	Id              string
	Type            string
	ServiceEndpoint string
	Revoked         bool
}

type DidDocument struct {
	Context   string
	Id        string
	Version   string
	PublicKey []*DidPublicKey
	Service   []*DidService
	Proof     proof.Proof
	Created   string
	Updated   string
	Revoked   bool
}

func (doc *DidDocument) AddPublicKey(pubKeys []*DidPublicKey) {
	if doc.PublicKey == nil {
		doc.PublicKey = pubKeys
		return
	}
	var existing bool
	for _, prev := range pubKeys {
		existing = false
		for _, element := range doc.PublicKey {
			if element.Id == prev.Id {
				existing = true
			}
		}
		if !existing {
			doc.PublicKey = append(doc.PublicKey, prev)
		}
	}
	return
}
func (doc *DidDocument) AddService(services []*DidService) {
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

func (doc *DidDocument) FindPublicKey(pubKeyId string) *DidPublicKey {
	if doc.PublicKey == nil || len(doc.PublicKey) == 0 {
		return nil
	}
	for _, pubKey := range doc.PublicKey {
		if pubKey.Id == pubKeyId {
			return pubKey
		}
	}
	return nil
}

func BuildDid(addressHex string) string {
	return fmt.Sprintf("did:pid:%s", addressHex)
}

func GetAddressFromDid(did string) string {
	//regExp := regexp.MustCompile("(?<=\bdid:pid:)[a-zA-Z0-9]*")
	ids := strings.Split(did, ":")
	if len(ids) == 3 {
		return ids[2]
	}
	return ""

}

func BuildPublicKeyId(addressHex string, idx int) string {
	return fmt.Sprintf("did:pid:%s#keys-%d", addressHex, idx)
}

func BuildDidPublicKeys(pubKeyId, pubKey, address string, pubKeyType PublicKeyType) string {
	publicKey := DidPublicKey{Id: pubKeyId, Type: string(pubKeyType), PublicKey: pubKey}
	jsonString, _ := json.Marshal([]DidPublicKey{publicKey})
	return string(jsonString)
}
