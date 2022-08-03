package did

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/keys/claimmeta"
	"github.com/datumtechs/did-sdk-go/keys/proof"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/datumtechs/did-sdk-go/types/algorithm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
	"time"
)

type VcService struct {
	ctx                chainclient.Context
	abi                abi.ABI
	vcContractInstance *contracts.Credential
	DocumentService    *DocumentService
	PctService         *PctService
}

func NewVcService(ctx chainclient.Context, documentService *DocumentService, pctService *PctService) *VcService {
	log.Info("Init Vc Service ...")
	m := new(VcService)
	m.ctx = ctx
	m.DocumentService = documentService
	m.PctService = pctService

	instance, err := contracts.NewCredential(vcContractAddress, ctx.GetClient())
	if err != nil {
		log.Fatal(err)
	}
	m.vcContractInstance = instance

	abiCode, err := abi.JSON(strings.NewReader(contracts.CredentialMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	m.abi = abiCode
	return m
}

type CreateCredentialReq struct {
	Context        string
	Type           string
	Issuer         string            // the issuer did
	PrivateKey     *ecdsa.PrivateKey // the private key to sign the credential
	PublicKeyId    string            // public key according to PublicKeyId should match the PrivateKey;if req.publicKeyId is provided, the issuer document should include the req.publicKeyId; else use the first valid public key in document.
	Did            string            // the applicant, vc holder
	PctId          *big.Int
	Claim          types.Claim
	ExpirationDate string
}

func (s *VcService) CreateCredentialSimple(req CreateCredentialReq) *Response[types.Credential] {
	return s.doCreateCredential(req, true)
}

func (s *VcService) CreateCredential(req CreateCredentialReq) *Response[types.Credential] {
	return s.doCreateCredential(req, false)
}

// check list:
// 1. req.Did, the applicant, vc holder should exist.
// 2. req.Issuer, the issuer Did should exist and valid
// 3. if req.publicKeyId is provided, the issuer document should include the req.publicKeyId; else use the first valid public key in document.
// 4. the req.privateKey and PublicKey in did document according to req.publicKeyId should be a pair of.
// 5. req.Claim should march the req.PctId's template
func (s *VcService) doCreateCredential(req CreateCredentialReq, simple bool) *Response[types.Credential] {
	// init the result
	response := new(Response[types.Credential])
	response.CallMode = false
	response.Status = Response_FAILURE

	if len(req.ExpirationDate) == 0 {
		req.ExpirationDate = common.FormatUTC(time.Now().AddDate(1, 0, 0).UTC())
	}
	//校验claim是否符合pctId所对应的json schema
	//5. req.Claim should march the req.PctId's template
	if !simple {
		verifyResp := s.PctService.VerifyByPct(req.PctId, req.Claim)
		if verifyResp.Status != Response_SUCCESS {
			CopyResp(verifyResp, response)
			return response
		}

		//1. req.Did, the applicant, vc holder should exist.
		address, err := types.ParseToAddress(req.Did)
		if err != nil {
			log.WithError(err).Errorf("failed to parse applicant did: %s", req.Did)
			response.Msg = "failed to parse applicant did"
			return response
		}

		checkDidResp := s.DocumentService.isDidExist(address)
		if checkDidResp.Status != Response_SUCCESS {
			CopyResp(checkDidResp, response)
			return response
		}
		if checkDidResp.Data == false {
			response.Msg = "did does not exist"
			return response
		}
		//2. req.Issuer, the issuer Did should exist and valid
		issuerAddress, err := types.ParseToAddress(req.Issuer)
		if err != nil {
			log.WithError(err).Errorf("failed to parse issuer did: %s", req.Issuer)

			response.Msg = "failed to parse issuer did"
			return response
		}

		checkIssuerDidResp := s.DocumentService.isDidExist(issuerAddress)
		if checkIssuerDidResp.Status != Response_SUCCESS {
			CopyResp(checkIssuerDidResp, response)
			return response
		}
		if checkIssuerDidResp.Data == false {
			response.Msg = "issuer did does not exist"
			return response
		}

		// check issuer did document status

		docStatusResp := s.DocumentService.GetDidDocumentStatus(issuerAddress)
		if docStatusResp.Status != Response_SUCCESS {
			CopyResp(docStatusResp, response)
			return response
		}
		if docStatusResp.Data != types.DOC_ACTIVATION {
			response.Msg = "issuer did document is DEACTIVATION"
			return response
		}

		// 3. the issuer document should include the req.publicKeyId
		// 4. the req.privateKey and PublicKey in did document according to req.publicKeyId should be a pair of.

		issuerDocResp := s.DocumentService.QueryDidDocumentByAddress(issuerAddress)
		if issuerDocResp.Status != Response_SUCCESS {
			CopyResp(issuerDocResp, response)
			response.Msg = "failed to find issuer did document"
			return response
		}
		issuerDoc := issuerDocResp.Data

		didPublicKeyToBeUsed := issuerDoc.FindDidPublicKeyByDidPublicKeyId(req.PublicKeyId)
		if didPublicKeyToBeUsed == nil {
			response.Msg = "failed to find public key ID to be used"
			return response
		} else if didPublicKeyToBeUsed.Status == types.PublicKey_INVALID.String() {
			response.Msg = "public key to be used is invalid"
			return response
		} else {
			pubKeyMatched := hex.EncodeToString(ethcrypto.FromECDSAPub(&req.PrivateKey.PublicKey))
			if didPublicKeyToBeUsed.PublicKey != pubKeyMatched {
				response.Msg = "public key in did document does not match to the private key"
				return response
			}
		}
	}
	// everything is ok
	//credentialWrapper := new(vc.CredentialWrapper)
	credential := generateCredential(req)

	digest := credential.GetDigest(nil, 0)

	//fmt.Printf("sign rawData: %s\n", rawData)
	sig := crypto.SignSecp256k1(digest, req.PrivateKey)

	//生成proof
	proofMap := make(types.Proof)
	proofMap[proofkeys.CREATED] = credential.IssuanceDate
	proofMap[proofkeys.TYPE] = algorithm.ALGO_SECP256K1
	proofMap[proofkeys.SIGNATURE] = hex.EncodeToString(sig)
	proofMap[proofkeys.VERIFICATIONMETHOD] = req.PublicKeyId
	credential.Proof = proofMap

	response.Data = *credential
	response.Status = Response_SUCCESS

	//todo: SaveProof on chain by another thread
	/*credentialWrapper.Credential = credential
	credentialWrapper.Disclosure = nil*/
	/*if !simple {
		go s.SaveVCProof(digestHash, s.ctx.GetAddress(), proofObj[proof.SIGNATURE])
	}*/

	response.Status = Response_SUCCESS
	return response
}

func generateCredential(req CreateCredentialReq) *types.Credential {
	//credentialWrapper := new(vc.CredentialWrapper)
	credential := new(types.Credential)
	credential.Id = uuid.NewString()
	credential.Context = req.Context
	if len(credential.Context) == 0 {
		credential.Context = types.DEFAULT_CREDENTIAL_CONTEXT
	}
	credential.Holder = req.Did
	credential.IssuanceDate = common.FormatUTC(time.Now().UTC())
	credential.ExpirationDate = req.ExpirationDate
	if len(credential.ExpirationDate) == 0 {
		req.ExpirationDate = common.FormatUTC(time.Now().AddDate(1, 0, 0).UTC())
	}
	credential.Version = types.VERSION
	credential.Type = []string{req.Type}
	credential.Issuer = req.Issuer

	credential.ClaimMeta = map[string]string{claimmetakeys.PCT_ID: req.PctId.String()}
	credential.ClaimData = req.Claim

	return credential
}

func (s *VcService) SaveVCProof(privateKey *ecdsa.PrivateKey, digestHash []byte, singer ethcommon.Address, proofSignature string) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = false
	response.Status = Response_FAILURE

	s.ctx.SetPrivateKey(privateKey)

	digestHash32 := ethcommon.BytesToHash(digestHash)

	// prepare parameters for createCredential()
	input, err := PackAbiInput(s.abi, "createCredential", digestHash32, singer.Hex(), proofSignature)
	if err != nil {
		log.WithError(err).Errorf("failed to pack input data for CreateCredential(), singer:%s", singer)
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, vcContractAddress, input)
	if err != nil {
		log.WithError(err).Errorf("failed to estimate gas for CreateCredential(), singer:%s", singer.Hex())
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)
	if err != nil {
		log.WithError(err).Errorf("failed to build tx options for CreateCredential(), singer:%s", singer.Hex())
		response.Msg = "failed to estimate gas"
		return response
	}

	// call contract CreatePid()
	tx, err := s.vcContractInstance.CreateCredential(opts, digestHash32, singer.Hex(), proofSignature)
	if err != nil {
		log.WithError(err).Errorf("failed to call CreateCredential(), singer:%s", singer.Hex())
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("call CreateCredential() txHash: %s", tx.Hash().Hex())

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
	response.Data = true
	return response
}

func (s *VcService) GetVCProof(credentialHash ethcommon.Hash) *Response[*types.ProofBrief] {
	// init the result
	response := new(Response[*types.ProofBrief])
	response.CallMode = true
	response.Status = Response_FAILURE

	blockNo, err := s.vcContractInstance.GetLatestBlock(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("failed to call VC GetLatestBlock(), credentialHash: %s", credentialHash.Hex())
		response.Msg = "failed to get latest block of VC proof"
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("VC proof not found, address: %s", credentialHash.Hex())
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
				response.Msg = "failed to parse contract event"
				return response
			}
			switch event.FieldKey {
			case types.VC_EVENT_SIGNATURE:
				proof.AddSignature(event.FieldValue)

				prevBlock = event.BlockNumber
			case types.VC_EVENT_SIGNERPUBKEY:
				proof.AddSignerPubKey(event.FieldValue)

				prevBlock = event.BlockNumber
			}
		}
	}

	response.Status = Response_SUCCESS
	response.Data = proof
	return response
}

func (s *VcService) HasVC(credentialHash ethcommon.Hash) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_FAILURE

	has, err := s.vcContractInstance.IsHashExist(nil, credentialHash)
	if err != nil {
		log.WithError(err).Errorf("failed to call IsHashExist(), error: %+v", err)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	response.Status = Response_SUCCESS
	response.Data = has
	return response
}

// 校验credential
// 首先，获取credential的指纹数据原文，以及credential的proof。signature；
// 然后获取issuer签发本vc用的public key；
// 最好做校验
// todo: 检查vc的claim是否符合pct定义；vc本身的状态; vc的有效期；检查issuer的签发公钥的状态
func (s *VcService) VerifyVC(credential *types.Credential) bool {

	// 从链上获取document
	dicDocResp := s.DocumentService.QueryDidDocument(credential.Issuer)
	if dicDocResp.Status == Response_SUCCESS {
		dicDoc := dicDocResp.Data
		didPubKey := dicDoc.FindDidPublicKeyByDidPublicKeyId(credential.Proof[proofkeys.VERIFICATIONMETHOD])
		if didPubKey == nil || len(didPubKey.PublicKey) == 0 {
			return false
		} else {
			return s.VerifyVCWithPublicKey(credential, crypto.HexToPublicKey(didPubKey.PublicKey))
			//return crypto.VerifySecp256k1Signature(rawData, credential.Proof[proof.SIGNATURE], crypto.HexToPublicKey(didPubKey.PublicKey))
		}
	} else {
		return false
	}
}

func (s *VcService) VerifyVCWithPublicKey(credential *types.Credential, publicKey *ecdsa.PublicKey) bool {
	sig := credential.Proof[proofkeys.SIGNATURE]

	digest := credential.GetDigest(nil, credential.ClaimData.GetSeed())
	//fmt.Printf("verify rawData: %s\n", rawData)

	return crypto.VerifySecp256k1Signature(digest, sig, publicKey)
}
