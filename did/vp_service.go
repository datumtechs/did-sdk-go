package did

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	proofkeys "github.com/datumtechs/did-sdk-go/keys/proof"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/datumtechs/did-sdk-go/types/algorithm"
	"time"
)

type DidAuthentication struct {
	Issuer           string            // the issuer did
	IssuerPrivateKey *ecdsa.PrivateKey // the private key to sign the presentation
	PublicKeyId      string            // public key identified by PublicKeyId in Did document should be consistent with PrivateKey; if req.publicKeyId is no provided, the first valid public key in Did document will be used.
}
type CreatePresentationReq struct {
	Authentication   DidAuthentication //签发VP的私钥等信息
	Credential       types.Credential
	Challenge        string
	DisclosurePolicy types.Claim //最终叶子节点的值，只能是0：不披露；1：披露
}

func (s *VcService) CreateVP(req CreatePresentationReq) *Response[types.Presentation] {
	//TODO: 检查req.DisclosurePolicy, 和req.Credential.ClaimMata定义的Pct是否一致。

	undisclosedSaltHash := common.Clone(req.Credential.ClaimData)
	seed := common.Uint64ToBigEndianBytes(req.Credential.ClaimData.GetSeed())

	types.SplitForMap(req.Credential.ClaimData, req.Credential.GetDisclosures(), undisclosedSaltHash, seed)

	presentation := types.Presentation{}
	presentation.VerifiableCredential = []types.Credential{req.Credential}
	digest := presentation.GetDigest()
	sig := crypto.SignSecp256k1(digest, req.Authentication.IssuerPrivateKey)

	//生成proof
	now := time.Now().UTC()
	createTime := common.FormatUTC(now)

	proofMap := make(types.Proof)
	proofMap[proofkeys.CREATED] = createTime
	proofMap[proofkeys.TYPE] = algorithm.ALGO_SECP256K1
	proofMap[proofkeys.JWS] = hex.EncodeToString(sig)
	proofMap[proofkeys.VERIFICATIONMETHOD] = req.Authentication.PublicKeyId
	presentation.Proof = proofMap
	return nil //s.doCreateCredential(did, context, pctId, claim, expirationDate, issuer, false)
}
