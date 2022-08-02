package did

import (
	"context"
	"crypto/ecdsa"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type DocumentService struct {
	ctx                 chainclient.Context
	abi                 abi.ABI
	didContractInstance *contracts.Did
}

func NewDocumentService(ctx chainclient.Context) *DocumentService {
	log.Info("Init Document service ...")
	m := new(DocumentService)
	m.ctx = ctx

	instance, err := contracts.NewDid(didContractAddress, ctx.GetClient())
	if err != nil {
		log.Fatal(err)
	}
	m.didContractInstance = instance

	abiCode, err := abi.JSON(strings.NewReader(contracts.DidABI))
	if err != nil {
		log.Fatal(err)
	}
	m.abi = abiCode
	return m
}

// datum 项目中，address就是carrier内置钱包地址。
type CreateDidReq struct {
	// Required: The private key of the signed transaction is also used to generate the corresponding public key, and the DID is generated from the public key.
	PrivateKey *ecdsa.PrivateKey
	// Required: The public key in DID Document.
	PublicKey string
	// Required: The type.  (default: Secp256k1)
	PublicKeyType types.PublicKeyType
}

func (s *DocumentService) CreateDID(req CreateDidReq) *Response[string] {
	// init the result
	response := new(Response[string])
	response.CallMode = false
	response.Status = Response_FAILURE

	s.ctx.SetPrivateKey(req.PrivateKey)

	// for did:pid:address
	address := ethcrypto.PubkeyToAddress(req.PrivateKey.PublicKey)
	did := types.BuildDid(address)

	// check if did exist
	didExistResp := s.isDidExist(address)
	if didExistResp.Status != Response_SUCCESS {
		CopyResp(didExistResp, response)
		return response
	}
	if didExistResp.Data == true {
		response.Msg = "Did exists already"
		return response
	}

	now := time.Now().UTC()
	createTime := common.FormatUTC(now)
	updateTime := createTime

	if len(req.PublicKeyType) == 0 {
		req.PublicKeyType = types.PublicKey_SECP256K1
	}
	publicKeyAsInput := types.BuildFieldValueOfPublicKey(req.PublicKey, req.PublicKeyType, "1", types.PublicKey_VALID)

	input, err := PackAbiInput(s.abi, "createDid", createTime, publicKeyAsInput, updateTime)
	if err != nil {
		log.WithError(err).Errorf("failed topack input data for CreateDid(),PublicKey:%s", req.PublicKey)
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, didContractAddress, input)
	if err != nil {
		log.WithError(err).Errorf("failed to estimate gas for CreateDid(),PublicKey:%s", req.PublicKey)
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)
	if err != nil {
		log.WithError(err).Errorf("failed to estimate gas for CreateDid(), PublicKey: %s", req.PublicKey)
		response.Msg = "failed to estimate gas"
		return response
	}

	// call contract CreateDid()
	tx, err := s.didContractInstance.CreateDid(opts, createTime, publicKeyAsInput, updateTime)
	if err != nil {
		log.WithError(err).Errorf("failed to call CreateDid(), PublicKey: %s", publicKeyAsInput)
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("call CreateDid() txHash:%s, PublicKey:%s", tx.Hash().Hex(), publicKeyAsInput)

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = Response_UNKNOWN
		response.Data = did
		response.Msg = "failed to get tx receipt"
		return response
	}
	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Msg = "failed to process tx"
		return response
	}
	// 交易信息
	response.TxInfo = NewTransactionInfo(receipt)
	// 返回did
	response.Data = did
	response.Status = Response_SUCCESS
	return response
}

func (s *DocumentService) QueryDidDocument(did string) *Response[*types.DidDocument] {
	// init the result
	response := new(Response[*types.DidDocument])
	response.CallMode = true
	response.Status = Response_FAILURE

	address, err := types.ParseToAddress(did)
	if err != nil {
		log.WithError(err).Errorf("failed to parse did: %s", did)
		response.Msg = "failed to parse did"
		return response
	}
	return s.QueryDidDocumentByAddress(address)
}

func (s *DocumentService) QueryDidDocumentByAddress(address ethcommon.Address) *Response[*types.DidDocument] {
	// init the result
	response := new(Response[*types.DidDocument])
	response.CallMode = true
	response.Status = Response_FAILURE

	blockNo, err := s.didContractInstance.GetLatestBlock(nil, address)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetLatestBlock(), address: %s", address)
		response.Msg = "failed to get latest block of DID"
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("DID not found, address: %s", address)
		response.Msg = "DID not found"
		return response
	}

	document := new(types.DidDocument)
	document.Id = types.BuildDid(address)

	// 查询document状态
	docStatusResp := s.GetDidDocumentStatus(address)
	if docStatusResp.Status != Response_SUCCESS {
		CopyResp(docStatusResp, response)
		return response
	}
	document.Status = docStatusResp.Data.String()

	// 遍历区块日志，查询document其它数据
	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	prevBlock := blockNo
	for prevBlock.Uint64() >= 0 {
		logs := s.ctx.GetLog(timeoutCtx, didContractAddress, prevBlock)
		for _, eachLog := range logs {
			event, err := s.didContractInstance.ParseDIDAttributeChange(*eachLog)
			if err != nil {
				response.Msg = "failed to parse contract event"
				return response
			}
			switch event.FieldKey {
			case types.DOC_EVENT_CREATE:
				document.Created = event.FieldValue
				document.Updated = event.UpdateTime
				prevBlock = event.BlockNumber
			case types.DOC_EVEN_PUBLICKEY:
				didPublicKey := types.ParseToDidPublicKey(document.Id, event.FieldValue)
				document.Updated = event.UpdateTime
				document.AddDidPublicKey(didPublicKey)
				prevBlock = event.BlockNumber
			}
		}
	}

	response.Data = document
	response.Status = Response_SUCCESS
	return response

}

type AddPublicKeyReq struct {
	// Required: The private key of the signed transaction is also used to generate the corresponding public key, and the DID is generated from the public key.
	PrivateKey *ecdsa.PrivateKey
	// Required: The public key in DID Document.
	PublicKey string
	// Required: The type.  (default: Secp256k1)
	PublicKeyType types.PublicKeyType
	//序号
	Index int
}

// 这是对合约setAttribute()方法的一个包装
func (s *DocumentService) AddPublicKey(req AddPublicKeyReq) *Response[bool] {
	response := new(Response[bool])
	response.CallMode = false

	response.Status = Response_SUCCESS

	s.ctx.SetPrivateKey(req.PrivateKey)

	// for did:pid:address
	address := ethcrypto.PubkeyToAddress(req.PrivateKey.PublicKey)
	did := types.BuildDid(address)
	// check if did exist
	didExistResp := s.isDidExist(address)
	if didExistResp.Status != Response_SUCCESS {
		CopyResp(didExistResp, response)
		return response
	}
	if didExistResp.Data == false {
		response.Status = Response_FAILURE
		response.Msg = "Did does not exist"
		return response
	}

	// did 存在，则查询document；如果getLatestBlock返回0或者-1表示不存在，则无需上面的检查
	docResp := s.QueryDidDocument(did)
	if docResp.Status != Response_SUCCESS {
		CopyResp(docResp, response)
		return response
	}
	didDoc := docResp.Data

	newPublicKeyId := types.BuildPublicKeyId(did, strconv.Itoa(req.Index))
	//to check if DID document has ths public key id already?
	if exist := didDoc.IsPublicKeyIdOrPublicKeyExist(newPublicKeyId, req.PublicKey); exist {
		log.Warningf("Public key or index exist: req:%+v", req)
		response.Status = Response_FAILURE
		response.Msg = "Public key or index exist"
		return response
	}

	now := time.Now().UTC()
	updateTime := common.FormatUTC(now)

	fieldValue := types.BuildFieldValueOfPublicKey(req.PublicKey, req.PublicKeyType, strconv.Itoa(req.Index), types.PublicKey_VALID)
	input, err := PackAbiInput(s.abi, "setAttribute", types.DOC_EVEN_PUBLICKEY, fieldValue, updateTime)

	if err != nil {
		log.Errorf("failed to pack input data for SetAttribute(): %+v", err)
		response.Status = Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, didContractAddress, input)
	if err != nil {
		log.Errorf("failed to estimate gas for SetAttribute(): %s, address: %v", address, err)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.didContractInstance.SetAttribute(opts, types.DOC_EVEN_PUBLICKEY, fieldValue, updateTime)
	if err != nil {
		log.WithError(err).Errorf("failed to call SetAttribute(), address: %s", address)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	log.Debugf("call SetAttribute() txHash:%s, addresss:%s", tx.Hash().Hex(), address)

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
	// 返回true
	response.Data = true
	response.Status = Response_SUCCESS
	return response
}

func (s *DocumentService) isDidExist(address ethcommon.Address) *Response[bool] {
	response := new(Response[bool])
	response.CallMode = true

	existing, err := s.didContractInstance.IsIdentityExist(nil, address)
	if err != nil {
		response.Msg = "failed to check if did exist"
		response.Status = Response_FAILURE
		return response
	}
	response.Data = existing
	response.Status = Response_SUCCESS
	return response
}

func (s *DocumentService) GetDidDocumentStatus(address ethcommon.Address) *Response[types.DocumentStatus] {
	response := new(Response[types.DocumentStatus])
	response.CallMode = true
	response.Status = Response_FAILURE

	status, err := s.didContractInstance.GetStatus(nil, address)
	if err != nil {
		log.WithError(err).Errorf("failed to find did document status,address:%s", address)
		response.Msg = "failed to find did document status"
		return response
	}
	//  -1：不存在
	if status == -1 {
		response.Msg = "did document does not exist"
		return response
	}
	response.Data = types.DocumentStatus(status)
	response.Status = Response_SUCCESS
	return response
}

/*func (s *DocumentService) HasPublicKey(address ethcommon.Address, pubKey string) *Response[bool] {
	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_SUCCESS

	if docResp := s.GetDocument(address); docResp.Status != Response_SUCCESS {
		response.Msg = docResp.Msg
		response.Status = docResp.Status
		return response
	} else {
		doc := docResp.Data
		if doc.FindDidPublicKeyByPublicKey(pubKey) != nil {
			response.Data = true
		} else {
			response.Data = false
		}
		return response
	}
}*/
/*
func (s *DocumentService) GetDidPublicKey(address ethcommon.Address, pubKey string) *Response[*doc.DidPublicKey] {
	response := new(Response[*doc.DidPublicKey])
	response.CallMode = true
	response.Status = Response_SUCCESS

	if docResp := s.QueryDidDocument(address); docResp.Status != Response_SUCCESS {
		response.Msg = docResp.Msg
		response.Status = docResp.Status
		return response
	} else {
		doc := docResp.Data
		if didPubKey := doc.FindDidPublicKeyByPublicKey(pubKey); didPubKey != nil {
			response.Data = didPubKey
			response.Status = Response_SUCCESS
		} else {
			response.Status = Response_FAILURE
			return response
		}
		return response
	}
}
*/
