package did

import (
	"context"
	"crypto/ecdsa"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/keys/proof"
	"github.com/datumtechs/did-sdk-go/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethhexuti "github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"time"
)

type CreateEvidenceReq struct {
	Credential types.Credential  // Required: the signed credential
	PrivateKey *ecdsa.PrivateKey `json:"-"` // Required: The private key of to sign transaction
}

func (s *CredentialService) CreateEvidence(req CreateEvidenceReq) *Response[string] {
	// init the result
	response := new(Response[string])
	response.CallMode = false
	response.Status = Response_FAILURE

	s.ctx.SetPrivateKey(req.PrivateKey)

	// 从链上获取document
	docResp := s.DocumentService.QueryDidDocument(req.Credential.Issuer)
	if docResp.Status != Response_SUCCESS {
		CopyResp(docResp, response)
		return response
	}
	checkDocResp := s.DocumentService.VerifyDocument(docResp.Data, req.Credential.Proof[proofkeys.VERIFICATIONMETHOD].(string), nil)
	if checkDocResp.Status != Response_SUCCESS {
		CopyResp(checkDocResp, response)
		return response
	}

	pubkeyHex := ethhexuti.Encode(ethcrypto.FromECDSAPub(checkDocResp.Data))

	// verify VC 签名
	ok := s.VerifyCredentialWithPublicKey(&req.Credential, checkDocResp.Data)
	if !ok {
		response.Msg = "failed to verify credential"
		return response
	}

	updateTime := common.FormatUTC(time.Now().UTC())

	seed, err := req.Credential.Proof.GetSeed()
	if err != nil {
		response.Msg = "failed to parse seed"
		return response
	}

	credentialHash, _ := req.Credential.GetHash(seed)

	// prepare parameters for createCredential()
	input, err := PackAbiInput(s.abi, "createCredential", credentialHash, pubkeyHex, req.Credential.Proof[proofkeys.JWS].(string), updateTime)
	if err != nil {
		log.WithError(err).Errorf("CreateEvidence: failed to pack input data, credential ID:%s", req.Credential.Id)
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.credentialContractProxy, input)
	if err != nil {
		log.WithError(err).Errorf("CreateEvidence: failed to estimate gas, credential ID:%s", req.Credential.Id)
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)
	if err != nil {
		log.WithError(err).Errorf("CreateEvidence: failed to build tx options, credential ID:%s", req.Credential.Id)
		response.Msg = "failed to build tx options"
		return response
	}

	// call contract CreatePid()
	tx, err := s.credentialContractInstance.CreateCredential(opts, credentialHash, pubkeyHex, req.Credential.Proof[proofkeys.JWS].(string), updateTime)
	if err != nil {
		log.WithError(err).Errorf("CreateEvidence: failed to call contract, credential ID:%s", req.Credential.Id)
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("CreateEvidence: call contract txHash: %s", tx.Hash().Hex())

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = Response_UNKNOWN
		response.Msg = "failed to get tx receipt"
	}
	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Msg = "failed to process tx"
	}

	// 交易信息
	response.TxInfo = NewTransactionInfo(receipt)
	response.Status = Response_SUCCESS
	response.Data = credentialHash.Hex()
	return response
}

type QueryEvidenceReq struct {
	EvidenceId string // credential digest, it is used to generate proof's signature by issuer
}

func (s *CredentialService) QueryEvidence(req QueryEvidenceReq) *Response[*types.EvidenceSignInfo] {
	// init the result
	response := new(Response[*types.EvidenceSignInfo])
	response.CallMode = true
	response.Status = Response_SUCCESS

	credentialHash := ethcommon.BytesToHash(ethhexuti.MustDecode(req.EvidenceId))

	status, err := s.credentialContractInstance.GetStatus(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("QueryEvidence: failed to get latest block), evidenceId: %s", req.EvidenceId)
		response.Msg = "failed to get latest block"
		response.Status = Response_FAILURE
		return response
	}
	// init return struct
	evidence := new(types.EvidenceSignInfo)
	evidence.CredentialHash = req.EvidenceId
	// 设置 credential 状态，也就是proof/evidence状态
	evidence.Status = types.CredentialStatus(status).String()

	blockNo, err := s.credentialContractInstance.GetLatestBlock(nil, ethcommon.BytesToHash(ethhexuti.MustDecode(req.EvidenceId)))
	if err != nil {
		log.WithError(err).Errorf("QueryEvidence: failed to get latest block), evidenceId: %s", req.EvidenceId)
		response.Msg = "failed to get latest block"
		response.Status = Response_FAILURE
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("QueryEvidence: evidence not found, evidenceId:%s", req.EvidenceId)
		response.Msg = "evidence not found"
		response.Status = Response_FAILURE
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	prevBlock := blockNo
	breakFor := false
	go func() {
	main_loop:
		for !breakFor && prevBlock.Uint64() > 0 {
			logs, err := s.ctx.GetLog(timeoutCtx, s.credentialContractProxy, prevBlock)
			if err != nil {
				log.WithError(err).Errorf("failed to get block logs, blockNumber:%d", prevBlock.Uint64())
				response.Msg = "failed to get block logs"
				response.Status = Response_FAILURE
				break main_loop
			}
			for _, eachLog := range logs {
				event, err := s.credentialContractInstance.ParseCredentialAttributeChange(eachLog)
				if err != nil {
					response.Msg = "failed to parse CredentialAttributeChange event"
					response.Status = Response_FAILURE
					break main_loop
				}
				switch event.FieldKey {
				case types.VC_EVENT_SIGNATURE:
					evidence.SupplementSignature(event.FieldValue)
					evidence.Timestamp = event.UpdateTime
				case types.VC_EVENT_SIGNERPUBKEY:
					evidence.SupplementSignerPubKey(event.FieldValue)
					evidence.Timestamp = event.UpdateTime
				}
				prevBlock = event.BlockNumber
			}
		}
		cancelFn()
	}()

	<-timeoutCtx.Done()
	breakFor = true

	response.Data = evidence
	return response
}

type RevokeEvidenceReq struct {
	EvidenceId string            // credential digest, it is used to generate proof's signature by issuer
	PrivateKey *ecdsa.PrivateKey `json:"-"` // Required: The private key of to sign transaction
}

func (s *CredentialService) RevokeEvidence(req RevokeEvidenceReq) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_FAILURE

	credentialHash := ethcommon.BytesToHash(ethhexuti.MustDecode(req.EvidenceId))

	status, err := s.credentialContractInstance.GetStatus(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("RevokeEvidence: failed to get latest block), evidenceId: %s", req.EvidenceId)
		response.Msg = "failed to get latest block"
		return response
	}
	if types.Credential_INVALID == types.CredentialStatus(status) {
		log.Warningf("RevokeEvidence: evidence is invalid already, evidenceId: %s", req.EvidenceId)
		response.Msg = "evidence is invalid already"
		return response
	}

	// init the tx signer
	s.ctx.SetPrivateKey(req.PrivateKey)

	// prepare parameters for EffectProposal()
	input, err := PackAbiInput(s.abi, "changeStatus", credentialHash, int8(types.Credential_INVALID))
	if err != nil {
		log.WithError(err).Errorf("RevokeEvidence: failed to pack input, evidenceId:%s", req.EvidenceId)
		response.Status = Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.credentialContractProxy, input)
	if err != nil {
		log.WithError(err).Errorf("RevokeEvidence: failed to estimate gas, evidenceId:%s", req.EvidenceId)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract ChangeStatus()
	tx, err := s.credentialContractInstance.ChangeStatus(opts, credentialHash, int8(types.Credential_INVALID))
	if err != nil {
		log.WithError(err).Errorf("RevokeEvidence:  failed to call contract, evidenceId:%s", req.EvidenceId)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("CreateEvidence: call contract txHash: %s", tx.Hash().Hex())

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = Response_UNKNOWN
		response.Msg = "failed to get tx receipt"
		return response
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Status = Response_FAILURE
		response.Msg = "failed to process tx"
		return response
	}
	// 交易信息
	response.TxInfo = NewTransactionInfo(receipt)
	response.Data = true
	response.Status = Response_SUCCESS
	return response
}

type VerifyCredentialEvidenceReq struct {
	Credential types.Credential
}

// VerifyCredentialEvidence 首先校验credential的proof是否符合credential；然后校验proof对应的evidence
func (s *CredentialService) VerifyCredentialEvidence(req VerifyCredentialEvidenceReq) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_FAILURE

	//查询issuer的document
	queryIssuerDocResp := s.DocumentService.QueryDidDocument(req.Credential.Issuer)
	if queryIssuerDocResp.Status != Response_SUCCESS {
		log.Errorf("VerifyCredentialEvidence: issuer did document not found, issuer: %s", req.Credential.Issuer)
		response.Msg = "issuer did document not found"
		return response
	}
	//查询issuer签发用publicKey
	issuerPublicKeyHex := queryIssuerDocResp.Data.FindDidPublicKeyByDidPublicKeyId(req.Credential.Proof[proofkeys.VERIFICATIONMETHOD].(string)).PublicKey
	issuerPublicKey := crypto.HexToPublicKey(issuerPublicKeyHex)
	if issuerPublicKey == nil {
		log.Errorf("VerifyCredentialEvidence: cannot unmarshal public key from did document: %s", issuerPublicKeyHex)
		response.Msg = "cannot unmarshal public key"
		return response
	}

	//查询evidence的状态

	seed, err := req.Credential.Proof.GetSeed()
	if err != nil {
		response.Msg = "failed to parse seed"
		return response
	}

	credentialHash, _ := req.Credential.GetHash(seed)

	status, err := s.credentialContractInstance.GetStatus(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("VerifyCredentialEvidence: failed to get latest block), evidenceId: %s", credentialHash.Hex())
		response.Msg = "failed to get latest block"
		return response
	}
	if types.Credential_INVALID == types.CredentialStatus(status) {
		log.Errorf("VerifyCredentialEvidence: evidence is invalid, evidenceId: %s", credentialHash.Hex())
		response.Msg = "evidence is invalid"
		return response
	}
	//查询evidence的内容
	queryEvidenceReq := QueryEvidenceReq{}
	queryEvidenceReq.EvidenceId = credentialHash.Hex()
	queryEvidenceResp := s.QueryEvidence(queryEvidenceReq)
	if queryEvidenceResp.Status != Response_SUCCESS {
		log.Errorf("VerifyCredentialEvidence: evidence not found, evidenceId: %s", credentialHash.Hex())
		response.Msg = "evidence not found"
		return response
	}

	//比较credential和evidence
	if req.Credential.Proof[proofkeys.JWS] != queryEvidenceResp.Data.Signature || issuerPublicKeyHex != queryEvidenceResp.Data.SignerPubKey {
		log.Errorf("VerifyCredentialEvidence: evidence is not consistent with credential: %s", credentialHash.Hex())
		response.Msg = "evidence is not consistent with credential"
		return response
	}

	ok := s.VerifyCredentialWithPublicKey(&req.Credential, issuerPublicKey)
	response.Data = ok
	response.Status = Response_SUCCESS
	return response
}
