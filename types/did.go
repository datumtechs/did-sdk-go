package types

import (
	"encoding/json"
	"fmt"
)

type PublicKeyType string

const (
	PublicKey_RSA       PublicKeyType = "RSA"
	PublicKey_SECP256K1 PublicKeyType = "Secp256k1"
)

type PublicKeyStatus string

const (
	PublicKey_VALID   PublicKeyStatus = "0"
	PublicKey_INVALID PublicKeyStatus = "1"
)

type AuthenticationStatus string

const (
	Authentication_VALID   AuthenticationStatus = "0"
	Authentication_INVALID AuthenticationStatus = "1"
)

type ServiceStatus string

const (
	Service_VALID   ServiceStatus = "0"
	Service_INVALID ServiceStatus = "1"
)

type DidPublicKey struct {
	Id           string
	Type         string
	PublicKeyHex string
	Controller   string
	Status       string
}

type DidAuthentication struct {
	PublicKeyHex string
	Controller   string
	Status       string
}

type DidService struct {
	Id              string
	Type            string
	ServiceEndpoint string
	Status          string
}

type DidDocument struct {
	Context        string
	Id             string
	PublicKey      []DidPublicKey
	Authentication []DidAuthentication
	Service        []DidService
	Create         string
	Updated        string
	Status         string
}

func (doc *DidDocument) AddAuthentication(auths []DidAuthentication) {
	if doc.Authentication == nil {
		doc.Authentication = auths
		return
	}
	var existing bool
	for _, prev := range auths {
		existing = false
		for _, element := range doc.Authentication {
			if element.PublicKeyHex == prev.PublicKeyHex {
				existing = true
			}
		}
		if !existing {
			doc.Authentication = append(doc.Authentication, prev)
		}
	}
	return
}

func (doc *DidDocument) AddPublicKey(pubKeys []DidPublicKey) {
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
func (doc *DidDocument) AddService(services []DidService) {
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

func (doc *DidDocument) HasPublicKey(pubKeyId string) bool {
	if doc.PublicKey == nil || len(doc.PublicKey) == 0 {
		return false
	}
	for _, element := range doc.PublicKey {
		if element.Id == pubKeyId {
			return true
		}
	}
	return false
}

func BuildPid(addressHex string) string {
	return fmt.Sprintf("did:pid:%s", addressHex)
}
func BuildPublicKeyId(addressHex string, idx int) string {
	return fmt.Sprintf("did:pid:%s#keys-%d", addressHex, idx)
}

func BuildController(addressHex string) string {
	return addressHex
}

func BuildDidAuthentications(pubKeyHex, addressHex string, status AuthenticationStatus) string {
	auth := DidAuthentication{PublicKeyHex: pubKeyHex, Controller: BuildController(addressHex), Status: string(status)}
	jsonString, _ := json.Marshal([]DidAuthentication{auth})
	return string(jsonString)
}

func BuildDidPublicKeys(pubKeyId string, pubKeyHex, addressHex string, pubKeyType PublicKeyType, status PublicKeyStatus) string {
	pubKey := DidPublicKey{Id: pubKeyId, PublicKeyHex: pubKeyHex, Controller: BuildController(addressHex), Type: string(pubKeyType), Status: string(status)}
	jsonString, _ := json.Marshal([]DidPublicKey{pubKey})
	return string(jsonString)
}
