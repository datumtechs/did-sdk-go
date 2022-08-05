package did

import (
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
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"math/big"
	"math/rand"
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
	PublicKeyId    string            // public key identified by PublicKeyId in Did document should be consistent with PrivateKey; if req.publicKeyId is no provided, the first valid public key in Did document will be used.
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
// 4. the req.privateKey and PublicKey identified by PublicKeyId in Did document should be a pair of.
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

		checkApplicantDidResp := s.DocumentService.isDidExist(address)
		if checkApplicantDidResp.Status != Response_SUCCESS {
			CopyResp(checkApplicantDidResp, response)
			return response
		}
		if checkApplicantDidResp.Data == false {
			response.Msg = "did does not exist"
			return response
		}

		//2. req.Issuer, the issuer Did document should exist and valid
		docResp := s.DocumentService.QueryDidDocument(req.Issuer)
		if docResp.Status != Response_SUCCESS {
			CopyResp(docResp, response)
			return response
		}
		checkDocResp := s.DocumentService.VerifyDocument(docResp.Data, req.PublicKeyId, req.PrivateKey)
		if checkDocResp.Status != Response_SUCCESS {
			CopyResp(docResp, checkDocResp)
			return response
		}
	}
	// everything is ok
	seed := rand.Uint64()

	//credentialWrapper := new(vc.CredentialWrapper)
	credential := generateCredential(req)

	digest, rootHash := credential.GetDigest(seed)

	//fmt.Printf("sign rawData: %s\n", rawData)
	sig := crypto.SignSecp256k1(digest, req.PrivateKey)

	//生成proof
	proofMap := make(types.Proof)
	proofMap[proofkeys.CREATED] = credential.IssuanceDate
	proofMap[proofkeys.TYPE] = algorithm.ALGO_SECP256K1
	proofMap[proofkeys.JWS] = hex.EncodeToString(sig)
	proofMap[proofkeys.VERIFICATIONMETHOD] = req.PublicKeyId
	proofMap[proofkeys.CLAIM_ROOT_HASH] = rootHash
	proofMap[proofkeys.SEED] = seed
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

// VerifyVC verify the proof signature
// 首先，获取credential的指纹数据原文，以及credential的proof。signature；
// 然后获取issuer签发本vc用的public key；
// 最好做校验
// todo: 检查vc的claim是否符合pct定义；vc本身的状态; vc的有效期；检查issuer的签发公钥的状态
func (s *VcService) VerifyVC(credential *types.Credential) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_FAILURE

	// 从链上获取document
	docResp := s.DocumentService.QueryDidDocument(credential.Issuer)
	if docResp.Status != Response_SUCCESS {
		CopyResp(docResp, response)
		return response
	}
	checkDocResp := s.DocumentService.VerifyDocument(docResp.Data, credential.Proof[proofkeys.VERIFICATIONMETHOD].(string), nil)
	if checkDocResp.Status != Response_SUCCESS {
		CopyResp(docResp, checkDocResp)
		return response
	}
	response.Data = s.VerifyVCWithPublicKey(credential, checkDocResp.Data)
	response.Status = Response_SUCCESS
	return response
}

func (s *VcService) VerifyVCWithPublicKey(credential *types.Credential, publicKey *ecdsa.PublicKey) bool {
	sig := credential.Proof[proofkeys.JWS].(string)

	digest, _ := credential.GetDigest(credential.Proof[proofkeys.SEED].(uint64))
	//fmt.Printf("verify rawData: %s\n", rawData)

	return crypto.VerifySecp256k1Signature(digest, sig, publicKey)
}
