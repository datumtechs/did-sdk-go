package did

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	platoncommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/datumtechs/chainclient"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/crypto"
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
	ctx                      chainclient.Context
	abi                      abi.ABI
	documentContractInstance *contracts.Did
	documentContractProxy    ethcommon.Address
}

func NewDocumentService(ctx chainclient.Context, config *Config) *DocumentService {
	log.Info("Init Document service ...")
	m := new(DocumentService)
	m.ctx = ctx
	m.documentContractProxy = config.DocumentContractProxy

	instance, err := contracts.NewDid(m.documentContractProxy, ctx.GetClient())
	if err != nil {
		log.Fatal(err)
	}
	m.documentContractInstance = instance

	abiCode, err := abi.JSON(strings.NewReader(contracts.DidMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	m.abi = abiCode
	return m
}

// datum 项目中，address就是carrier内置钱包地址。
type CreateDidReq struct {
	// Required: The private key of the signed transaction is also used to generate the corresponding public key, and the DID is generated from the public key.
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	// Required: The public key in DID Document.
	PublicKey string
	// Required: The type.  (default: Secp256k1)
	PublicKeyType types.PublicKeyType
}

func (s *DocumentService) CreateDID(req CreateDidReq) *Response[string] {
	log.Debugf("CreateDID: req.PublicKey:%s", req.PublicKey)

	// init the result
	response := new(Response[string])
	response.CallMode = false
	response.Status = Response_FAILURE

	s.ctx.SetPrivateKey(req.PrivateKey)

	// for did:pid:bech32Addr
	address := ethcrypto.PubkeyToAddress(req.PrivateKey.PublicKey)
	did := types.BuildDid(address)
	response.Data = did

	// check if did exist
	didExistResp := s.isDidExist(address)
	if didExistResp.Status != Response_SUCCESS {
		CopyResp(didExistResp, response)
		return response
	}
	if didExistResp.Data == true {
		log.Errorf("Did exists, did: %s", did)
		response.Msg = "Did exists"
		response.Status = Response_EXIST
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
		log.WithError(err).Errorf("CreateDID: failed topack input data,PublicKey:%s", req.PublicKey)
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.documentContractProxy, input)
	if err != nil {
		log.WithError(err).Error("CreateDID: failed to estimate gas")
		response.Msg = err.Error()
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)
	if err != nil {
		log.WithError(err).Error("CreateDID: failed to build TxOpts")
		response.Msg = "failed to build TxOpts"
		return response
	}

	// call contract CreateDid()
	tx, err := s.documentContractInstance.CreateDid(opts, createTime, publicKeyAsInput, updateTime)
	if err != nil {
		log.WithError(err).Errorf("CreateDID: failed to call contrat, PublicKey: %s", publicKeyAsInput)
		response.Msg = err.Error()
		return response
	}
	log.Debugf("CreateDID: call contract, txHash:%s, PublicKey:%s", tx.Hash().Hex(), publicKeyAsInput)

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
	response.Status = Response_SUCCESS

	blockNo, err := s.documentContractInstance.GetLatestBlock(nil, address)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetLatestBlock(), bech32Addr: %s", platoncommon.Address(address).Bech32())
		response.Msg = "failed to get latest block of DID"
		response.Status = Response_FAILURE
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("DID not found, bech32Addr: %s", platoncommon.Address(address).Bech32())
		response.Msg = "DID not found"
		response.Status = Response_NOT_FOUND
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
	document.Status = docStatusResp.Data

	// 遍历区块日志，查询document其它数据
	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	prevBlock := blockNo
	log.Debugf("blockNumber:%d", prevBlock.Uint64())
	breakFor := false
	go func() {
	main_loop:
		for !breakFor && prevBlock.Uint64() > 0 {
			logs, err := s.ctx.GetLog(timeoutCtx, s.documentContractProxy, prevBlock)
			if err != nil {
				log.WithError(err).Errorf("failed to get block logs, blockNumber:%d", prevBlock.Uint64())
				response.Msg = "failed to get block logs"
				response.Status = Response_FAILURE
				break main_loop
			}
			for _, eachLog := range logs {
				event, err := s.documentContractInstance.ParseDIDAttributeChange(eachLog)
				if err != nil {
					response.Msg = "failed to parse contract event"
					response.Status = Response_FAILURE
					break main_loop
				}
				switch event.FieldKey {
				case types.DOC_EVENT_CREATE:
					document.Created = event.FieldValue
					document.Updated = event.UpdateTime
					prevBlock = event.BlockNumber
				case types.DOC_EVEN_PUBLICKEY:
					didPublicKey := types.EventToDidPublicKey(document.Id, event.FieldValue)
					document.Updated = event.UpdateTime
					document.SupplementDidPublicKey(didPublicKey)
					prevBlock = event.BlockNumber
				}
			}
		}
		cancelFn()
	}()

	<-timeoutCtx.Done()
	breakFor = true

	response.Data = document
	return response

}

type AddPublicKeyReq struct {
	// Required: The private key of the signed transaction is also used to generate the corresponding public key, and the DID is generated from the public key.
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	// Required: The public key in DID Document.
	PublicKey string
	// Required: The type.  (default: Secp256k1)
	PublicKeyType types.PublicKeyType
	//序号
	Index int
}

// 这是对合约setAttribute()方法的一个包装
func (s *DocumentService) AddPublicKey(req AddPublicKeyReq) *Response[bool] {
	log.Debugf("AddPublicKey: req.PublicKey:%s", req.PublicKey)
	response := new(Response[bool])
	response.CallMode = false

	response.Status = Response_SUCCESS

	s.ctx.SetPrivateKey(req.PrivateKey)

	// for did:pid:bech32Addr
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
		log.Warningf("AddPublicKey: Public key or index exist: req:%+v", req)
		response.Status = Response_EXIST
		response.Msg = "Public key or index exist"
		return response
	}

	now := time.Now().UTC()
	updateTime := common.FormatUTC(now)

	fieldValue := types.BuildFieldValueOfPublicKey(req.PublicKey, req.PublicKeyType, strconv.Itoa(req.Index), types.PublicKey_VALID)
	input, err := PackAbiInput(s.abi, "setAttribute", types.DOC_EVEN_PUBLICKEY, fieldValue, updateTime)

	if err != nil {
		log.WithError(err).Error("AddPublicKey: failed to pack input data")
		response.Status = Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.documentContractProxy, input)
	if err != nil {
		log.WithError(err).Error("AddPublicKey: failed to estimate gas")
		response.Status = Response_FAILURE
		response.Msg = err.Error()
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)
	if err != nil {
		log.WithError(err).Error("AddPublicKey: failed to build TxOpts")
		response.Msg = "failed to build TxOpts"
		return response
	}

	// call contract SetAttribute()
	tx, err := s.documentContractInstance.SetAttribute(opts, types.DOC_EVEN_PUBLICKEY, fieldValue, updateTime)
	if err != nil {
		log.WithError(err).Errorf("AddPublicKey: failed to call contract, address: %s", address)
		response.Status = Response_FAILURE
		response.Msg = err.Error()
		return response
	}

	log.Debugf("AddPublicKey: call contract, txHash:%s, address:%s", tx.Hash().Hex(), address)

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

	existing, err := s.documentContractInstance.IsIdentityExist(nil, address)
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

	status, err := s.documentContractInstance.GetStatus(nil, address)
	if err != nil {
		log.WithError(err).Errorf("failed to find did document status, address:%s", address.Hex())
		response.Msg = "failed to find did document status"
		return response
	}
	//  -1：不存在
	if status == -1 {
		response.Msg = "did document does not exist"
		response.Status = Response_NOT_FOUND
		return response
	}
	response.Data = types.DocumentStatus(status)
	response.Status = Response_SUCCESS
	return response
}

func (s *DocumentService) VerifyDocument(document *types.DidDocument, publicKeyId string, privateKey *ecdsa.PrivateKey) *Response[*ecdsa.PublicKey] {
	response := new(Response[*ecdsa.PublicKey])
	response.CallMode = true
	response.Status = Response_FAILURE

	if document == nil {
		response.Msg = "Did document not found"
		response.Status = Response_NOT_FOUND
		return response
	}
	if document.Status == types.DOC_DEACTIVATION {
		response.Msg = "Did document is DEACTIVATION"
		response.Status = Response_DEACTIVATION
		return response
	}
	didPublicKey := document.FindDidPublicKeyByDidPublicKeyId(publicKeyId)
	if didPublicKey == nil {
		response.Msg = "public key ID not found in Did document"
		response.Status = Response_NOT_FOUND
		return response
	}

	var pubkeyHex string = ""
	if didPublicKey.Status == types.PublicKey_INVALID {
		response.Msg = "The public key corresponding to the public key ID is INVALID"
		response.Status = Response_DEACTIVATION
		return response
	}
	pubkeyHex = didPublicKey.PublicKey
	if len(pubkeyHex) == 0 {
		response.Msg = "The public key corresponding to the public key ID is EMPTY"
		return response
	}

	if privateKey != nil {
		pubkeyBytesExpected := ethcrypto.FromECDSAPub(&privateKey.PublicKey)
		if !bytes.Equal(ethcommon.FromHex(pubkeyHex), pubkeyBytesExpected) {
			response.Msg = "public key in did document not consistent with the private key"
			return response
		}
	}
	response.Data = crypto.HexToPublicKey(pubkeyHex)
	response.Status = Response_SUCCESS
	return response
}

/*func (s *DocumentService) HasPublicKey(bech32Addr ethcommon.Address, pubKey string) *Response[bool] {
	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_SUCCESS

	if docResp := s.GetDocument(bech32Addr); docResp.Status != Response_SUCCESS {
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
func (s *DocumentService) GetDidPublicKey(bech32Addr ethcommon.Address, pubKey string) *Response[*doc.DidPublicKey] {
	response := new(Response[*doc.DidPublicKey])
	response.CallMode = true
	response.Status = Response_SUCCESS

	if docResp := s.QueryDidDocument(bech32Addr); docResp.Status != Response_SUCCESS {
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
