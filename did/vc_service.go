package did

import (
	"context"
	"encoding/hex"
	common2 "github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	vc_EVENT_FIELD_SIGNERPUBKEY uint8 = 0
	vc_EVENT_FIELD_SIGNATURE    uint8 = 1

	DEFAULT_CREDENTIAL_CONTEXT string = "https://www.w3.org/2018/credentials/v1"
)

func (s *DIDService) CreateCredential(did string, context string, pctId int, claim map[string]string, expirationDate string, issuer string) *common2.Response {
	// init the result
	response := new(common2.Response)
	response.CallMode = false

	credentialWrapper := new(types.CredentialWrapper)
	credential := new(types.Credential)

	if len(context) == 0 {
		context = DEFAULT_CREDENTIAL_CONTEXT
	}
	credential.Context = context
	credential.Id = uuid.NewString()
	credential.PctId = pctId
	credential.Issuer = issuer
	credential.IssuanceDate = common2.FormatUTC(time.Now().UTC())
	credential.ExpirationDate = expirationDate
	credential.Claim = claim
	credential.Holder = did

	//生成VC，默认claim所有字段都需要披露
	disclosureMap := make(map[string]int)
	for key, _ := range claim {
		disclosureMap[key] = types.CLAIM_FIELD_DISCLOSED
	}
	//生成
	credential.Proof = s.buildProof(credential, disclosureMap)

	credentialWrapper.Credential = credential
	credentialWrapper.Disclosure = disclosureMap

	response.Data = credentialWrapper
	return response

}

func (s *DIDService) buildProof(credential *types.Credential, disclosureMap map[string]int) map[string]interface{} {
	proof := make(map[string]interface{})
	proof["created"] = credential.IssuanceDate
	proof["creator"] = credential.Issuer
	proof["type"] = "Secp256k1"
	proof["signature"] = s.signCredential(credential, disclosureMap)
	return proof
}

func (did *DIDService) signCredential(credential *types.Credential, disclosureMap map[string]int) string {
	rawData := credential.GetCredentialThumbprintWithoutSig(disclosureMap)

	digestHash := crypto.Keccak256([]byte(rawData))
	sig, err := crypto.Sign(digestHash, did.ctx.GetPrivateKey())
	if err != nil {
		log.Errorf("failed to sign credential, error: %+v", err)
		return ""
	}
	return hex.EncodeToString(sig)
}

func (s *DIDService) SaveVCProof(credentialHash common.Hash, signerPubKey string, signature string) *common2.Response {
	// init the result
	response := new(common2.Response)
	response.CallMode = false

	// prepare parameters for submitProposal()
	input, err := s.packInput("createCredential", credentialHash, signerPubKey, signature)
	if err != nil {
		log.Errorf("failed to pack input data for CreateCredential(), error: %+v", err)
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to pack input data"
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, vcContractAddress, input)
	if err != nil {
		log.Errorf("failed to estimate gas for CreateCredential(), error: %+v", err)
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to estimate gas"
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.vcContractInstance.CreateCredential(opts, credentialHash, signerPubKey, signature)
	if err != nil {
		log.WithError(err).Errorf("failed to call CreateCredential(), error: %+v", err)
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to call contract"
	}
	response.TxHash = tx.Hash()
	response.Status = common2.Response_SUCCESS

	log.Debugf("call CreateCredential() txHash: %s", tx.Hash().Hex())

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = common2.Response_UNKNOWN
		response.Msg = "failed to get tx receipt"
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to process tx"
	} else {
		response.Status = common2.Response_SUCCESS
	}

	return response
}

func (s *DIDService) GetVCProof(credentialHash common.Hash) *common2.Response {
	// init the result
	response := new(common2.Response)
	response.CallMode = true

	blockNo, err := s.vcContractInstance.GetLatestBlock(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("failed to call VC GetLatestBlock(), credentialHash: %s", credentialHash.Hex())
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to get latest block of VC proof"
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("VC proof not found, address: %s", credentialHash.Hex())
		response.Status = common2.Response_FAILURE
		response.Msg = "VC proof not found"
		return response
	}

	proof := new(types.ProofBrief)
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
				response.Status = common2.Response_FAILURE
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

func (s *DIDService) HasVC(credentialHash common.Hash) *common2.Response {
	// init the result
	response := new(common2.Response)
	response.CallMode = true

	has, err := s.vcContractInstance.IsHashExist(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("failed to call IsHashExist(), error: %+v", err)
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	response.Data = has
	return response
}
