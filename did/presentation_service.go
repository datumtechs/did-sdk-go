package did

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	claimmetakeys "github.com/datumtechs/did-sdk-go/keys/claimmeta"
	proofkeys "github.com/datumtechs/did-sdk-go/keys/proof"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/datumtechs/did-sdk-go/types/algorithm"
	"time"
)

type Authentication struct {
	Issuer           string            // the issuer did
	IssuerPrivateKey *ecdsa.PrivateKey // the private key to sign the presentation
	PublicKeyId      string            // public key identified by PublicKeyId in Did document should be consistent with PrivateKey; if req.publicKeyId is no provided, the first valid public key in Did document will be used.
}
type CreatePresentationReq struct {
	Authentication     Authentication      //签发VP的私钥等信息
	Credential         []*types.Credential //必须是指针类型
	Challenge          string
	PresentationPolicy map[string]types.Claim //key is pct_Id
}

func (s *VcService) CreateVP(req CreatePresentationReq) *Response[types.Presentation] {
	// init the result
	response := new(Response[types.Presentation])
	response.CallMode = false
	response.Status = Response_FAILURE

	//TODO: 检查req.DisclosurePolicy, 和req.Credential.ClaimMata定义的Pct是否一致。
	for idx := 0; idx < len(req.Credential); idx++ { //range 会copy，用下标访问
		//for _, credential := range req.Credential {
		// 当req.Credential是个struct的slice，如： []types.Credential
		// 赋值必须是：credential := &req.Credential[idx]
		// 否则credential的作用域只是在for{}中；
		// 总结：变量credential必须是指针。
		// 另外：当credential中的变量不是基础类型，或者struct时，如map,slice等时，不管credential是否是指针，对credential.field，实际上都是指向原始地址，在for{}中的修改会生效。{}
		// credential := &req.Credential[idx]
		credential := req.Credential[idx]
		disclosures := req.PresentationPolicy[credential.ClaimMeta[claimmetakeys.PCT_ID]]
		credential.Proof[proofkeys.DISCLOSURES] = disclosures
		seed := common.Uint64ToBigEndianBytes(credential.Proof[proofkeys.SEED].(uint64))
		//这里credential.ClaimData是个map，传第的是个指针，所以函数内部的修改会生效
		err := types.SplitForMap(credential.ClaimData, disclosures, seed)
		if err != nil {
			response.Msg = "failed to analyse clmain and it's disclosures"
			return response
		}

	}
	presentation := types.Presentation{}
	presentation.VerifiableCredential = req.Credential
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
	proofMap[proofkeys.CHALLENGE] = req.Challenge
	presentation.Proof = proofMap

	response.Status = Response_SUCCESS
	response.Data = presentation
	return response
}

type VerifyPresentationReq struct {
	Challenge          string
	Holder             string // vp holder's DID
	Presentation       types.Presentation
	PresentationPolicy map[string]types.Claim //key is pct_Id
}

func (s *VcService) VerifyVP(req VerifyPresentationReq) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = false
	response.Status = Response_FAILURE

	// step1: to check each presentation
	// find the vp holder's did document
	docResp := s.DocumentService.QueryDidDocument(req.Holder)
	if docResp.Status != Response_SUCCESS {
		CopyResp(docResp, response)
		return response
	}
	checkDocResp := s.DocumentService.VerifyDocument(docResp.Data, req.Presentation.Proof[proofkeys.VERIFICATIONMETHOD].(string), nil)
	if checkDocResp.Status != Response_SUCCESS {
		CopyResp(docResp, checkDocResp)
		return response
	}

	ok := crypto.VerifySecp256k1Signature(req.Presentation.GetDigest(), req.Presentation.Proof[proofkeys.JWS].(string), checkDocResp.Data)
	if !ok {
		response.Msg = "failed to verify presentation proof"
		return response
	}

	// step2: to check each credential

	for idx := 0; idx < len(req.Presentation.VerifiableCredential); idx++ {
		credential := req.Presentation.VerifiableCredential[idx]
		/*		claimPolicy := credential.Proof[proofkeys.DISCLOSURES].(types.Claim)
				seed := credential.Proof[proofkeys.SEED].(uint64)
				claimRootHash := credential.Proof[proofkeys.CLAIM_ROOT_HASH].(string)
				pubkeyId := credential.Proof[proofkeys.VERIFICATIONMETHOD].(string)
		*/
		checkVcResp := s.VerifyVC(credential)
		if checkVcResp.Status != Response_SUCCESS {
			CopyResp(checkVcResp, response)
			return response
		}
		if !checkVcResp.Data {
			response.Msg = "failed to verify credential"
			return response
		}
	}
	response.Status = Response_SUCCESS
	response.Data = true
	return response
}
