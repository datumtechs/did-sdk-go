package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

func SignSecp256k1(rawData string, privateKey *ecdsa.PrivateKey) (digestHash []byte, signature string) {
	//digestHash := ethcrypto.Keccak256([]byte(rawData))
	digestHash = SHA3(rawData)
	sig, err := ethcrypto.Sign(digestHash, privateKey)
	if err != nil {
		log.Errorf("failed to sign credential, error: %+v", err)
		return nil, ""
	}
	signature = hex.EncodeToString(sig)
	return
}

func VerifySecp256k1Signature(rawData string, signature string, publicKey *ecdsa.PublicKey) bool {
	if len(rawData) == 0 || len(signature) == 0 || publicKey == nil {
		return false
	}
	if signature, err := hex.DecodeString(signature); err != nil {
		log.Errorf("failed to decode signature hex string, error: %+v", err)
		return false
	} else {
		// remove recovery id (signature[64]
		//return ethcrypto.VerifySignature(ethcrypto.FromECDSAPub(publicKey), ethcrypto.Keccak256([]byte(rawData)), signature[:len(signature)-1])
		return ethcrypto.VerifySignature(ethcrypto.FromECDSAPub(publicKey), SHA3(rawData), signature[:len(signature)-1])
	}
}

func HexToPublicKey(publicKey string) *ecdsa.PublicKey {
	if pubKey, err := hex.DecodeString(publicKey); err == nil {
		if pk, err := ethcrypto.UnmarshalPubkey(pubKey); err == nil {
			return pk
		} else {
			log.Errorf("failed to unmarshal public key, error: %+v", err)
		}
	} else {
		log.Errorf("failed to decode hex public key , error: %+v", err)
	}
	return nil
}
