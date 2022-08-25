package crypto

import (
	"crypto/ecdsa"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethhexutl "github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

func SignSecp256k1(digest []byte, privateKey *ecdsa.PrivateKey) []byte {
	//digestHash := ethcrypto.Keccak256([]byte(rawData))
	//digestHash = LegacyKeccak256SHA3(rawData)
	sig, err := ethcrypto.Sign(digest, privateKey)
	if err != nil {
		log.WithError(err).Errorf("failed to sign credential")
		return nil
	}
	return sig
}

func VerifySecp256k1Signature(digest []byte, signature string, publicKey *ecdsa.PublicKey) bool {
	if len(digest) == 0 || len(signature) == 0 || publicKey == nil {
		return false
	}
	if signature, err := ethhexutl.Decode(signature); err != nil {
		log.WithError(err).Errorf("failed to decode signature hex:%s", signature)
		return false
	} else {
		// remove recovery id (signature[64]
		//return ethcrypto.VerifySignature(ethcrypto.FromECDSAPub(publicKey), ethcrypto.Keccak256([]byte(rawData)), signature[:len(signature)-1])
		return ethcrypto.VerifySignature(ethcrypto.FromECDSAPub(publicKey), digest, signature[:len(signature)-1])
	}
}

// HexToPublicKey decodes a hex string with 0x prefix as an ecdsa.PublicKey.
func HexToPublicKey(publicKey string) *ecdsa.PublicKey {
	pubKey := ethcommon.FromHex(publicKey)
	if pk, err := ethcrypto.UnmarshalPubkey(pubKey); err == nil {
		return pk
	} else {
		log.WithError(err).Errorf("failed to unmarshal public key:%s", publicKey)
	}
	return nil
}
