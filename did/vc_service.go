package did

import (
	"context"
	"crypto/ecdsa"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/types/claim"
	"github.com/datumtechs/did-sdk-go/types/doc"
	"github.com/datumtechs/did-sdk-go/types/proof"
	"github.com/datumtechs/did-sdk-go/types/vc"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	vc_EVENT_FIELD_SIGNERPUBKEY uint8 = 0
	vc_EVENT_FIELD_SIGNATURE    uint8 = 1

	DEFAULT_CREDENTIAL_CONTEXT string = "https://www.w3.org/2018/credentials/v1"
)

func (s *DIDService) CreateCredential(did string, context string, pctId int, claim claim.Claim, expirationDate string, issuer string) *Response[vc.CredentialWrapper] {
	// init the result
	response := new(Response[vc.CredentialWrapper])
	response.CallMode = false

	credentialWrapper := new(vc.CredentialWrapper)
	credential := new(vc.Credential)

	if len(context) == 0 {
		context = DEFAULT_CREDENTIAL_CONTEXT
	}
	credential.Context = context
	credential.Id = uuid.NewString()
	credential.PctId = pctId
	credential.Issuer = issuer
	credential.IssuanceDate = common.FormatUTC(time.Now().UTC())
	credential.ExpirationDate = expirationDate
	credential.Claim = claim
	credential.Holder = did

	//生成
	credential.Proof = s.BuildProof(credential, nil)

	credentialWrapper.Credential = credential
	credentialWrapper.Disclosure = nil

	response.Data = *credentialWrapper
	response.Status = Response_SUCCESS
	return response

}

func (s *DIDService) BuildProof(credential *vc.Credential, disclosureMap map[string]int) proof.Proof {
	p := make(proof.Proof)
	p[proof.CREATED] = credential.IssuanceDate
	p[proof.TYPE] = "Secp256k1"
	p[proof.SIGNATURE] = s.SignCredential(credential, disclosureMap)
	return p
}

func (s *DIDService) SignCredential(credential *vc.Credential, disclosureMap map[string]int) string {
	rawData := credential.GetCredentialThumbprintWithoutSig(disclosureMap)
	return crypto.SignSecp256k1(rawData, s.ctx.GetPrivateKey())
}

func (s *DIDService) SaveVCProof(credentialHash ethcommon.Hash, signerPubKey string, signature string) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = false
	response.Data = false

	// prepare parameters for submitProposal()
	input, err := s.packInput("createCredential", credentialHash, signerPubKey, signature)
	if err != nil {
		log.Errorf("failed to pack input data for CreateCredential(), error: %+v", err)
		response.Status = Response_FAILURE
		response.Msg = "failed to pack input data"
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, vcContractAddress, input)
	if err != nil {
		log.Errorf("failed to estimate gas for CreateCredential(), error: %+v", err)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.vcContractInstance.CreateCredential(opts, credentialHash, signerPubKey, signature)
	if err != nil {
		log.WithError(err).Errorf("failed to call CreateCredential(), error: %+v", err)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
	}
	response.TxHash = tx.Hash()
	response.Status = Response_SUCCESS

	log.Debugf("call CreateCredential() txHash: %s", tx.Hash().Hex())

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = Response_UNKNOWN
		response.Msg = "failed to get tx receipt"
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Status = Response_FAILURE
		response.Msg = "failed to process tx"
	} else {
		response.Status = Response_SUCCESS
		response.Data = true
	}

	return response
}

func (s *DIDService) GetVCProof(credentialHash ethcommon.Hash) *Response[*vc.ProofBrief] {
	// init the result
	response := new(Response[*vc.ProofBrief])
	response.CallMode = true

	blockNo, err := s.vcContractInstance.GetLatestBlock(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("failed to call VC GetLatestBlock(), credentialHash: %s", credentialHash.Hex())
		response.Status = Response_FAILURE
		response.Msg = "failed to get latest block of VC proof"
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("VC proof not found, address: %s", credentialHash.Hex())
		response.Status = Response_FAILURE
		response.Msg = "VC proof not found"
		return response
	}

	proof := new(vc.ProofBrief)
	proof.CredentialHash = credentialHash

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	prevBlock := blockNo

	for prevBlock.Uint64() > 0 {
		logs := s.ctx.GetLog(timeoutCtx, vcContractAddress, prevBlock)
		for _, eachLog := range logs {
			event, err := s.vcContractInstance.ParseCredentialAttributeChange(*eachLog)
			if err != nil {
				response.Status = Response_FAILURE
				response.Msg = "failed to parse contract event"
				return response
			}
			switch event.FieldKey {
			case vc_EVENT_FIELD_SIGNATURE:
				proof.AddSignature(event.FieldValue)

				prevBlock = event.BlockNumber
			case vc_EVENT_FIELD_SIGNERPUBKEY:
				proof.AddSignerPubKey(event.FieldValue)

				prevBlock = event.BlockNumber
			}
		}
	}
	response.Data = proof
	return response
}

func (s *DIDService) HasVC(credentialHash ethcommon.Hash) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = true

	has, err := s.vcContractInstance.IsHashExist(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("failed to call IsHashExist(), error: %+v", err)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	response.Data = has
	return response
}

// 校验credential
// 首先，获取credential的指纹数据原文，以及credential的proof。signature；
// 然后获取issuer签发本vc用的public key；
// 最好做校验
// todo: 检查vc的claim是否符合pct定义；vc本身的状态; vc的有效期；检查issuer的签发公钥的状态
func (s *DIDService) VerifyCredential(credential *vc.Credential, disclosureMap map[string]int, publicKey *ecdsa.PublicKey) bool {
	rawData := credential.GetCredentialThumbprintWithoutSig(disclosureMap)
	issuerAddr := doc.GetAddressFromDid(credential.Issuer)

	if publicKey == nil {
		// 从链上获取document
		resp := s.GetDocument(ethcommon.HexToAddress(issuerAddr))
		if resp.Status == Response_SUCCESS {
			dicDoc := resp.Data
			didPubKey := dicDoc.FindPublicKey(credential.Proof[proof.PUBLIC_KEY_ID])
			if didPubKey == nil || len(didPubKey.PublicKey) == 0 {
				return false
			} else {
				return crypto.VerifySecp256k1Signature(rawData, credential.Proof[proof.SIGNATURE], crypto.HexToPublicKey(didPubKey.PublicKey))
			}
		} else {
			return false
		}
	} else {
		return crypto.VerifySecp256k1Signature(rawData, credential.Proof[proof.SIGNATURE], publicKey)
	}
}
