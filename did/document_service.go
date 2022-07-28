package did

import (
	"context"
	"encoding/json"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/types/doc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
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
func (s *DocumentService) CreateDID(address ethcommon.Address, pubKeyHex string) *Response[string] {
	// init the result
	response := new(Response[string])
	response.CallMode = false

	response.Status = Response_SUCCESS

	// prepare parameters for CreatePid()
	addrHex := address.Hex()
	now := time.Now().UTC()
	createTime := common.FormatUTC(now)
	updateTime := createTime

	pubKeyId := doc.BuildPublicKeyId(addrHex, 1)
	publicKey := doc.BuildDidPublicKeys(pubKeyId, pubKeyHex, addrHex, doc.PublicKey_SECP256K1)

	input, err := PackAbiInput(s.abi, "createPid", createTime, publicKey, updateTime)

	if err != nil {
		log.Errorf("failed to pack input data for CreatePid(): %+v", err)
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
		log.Errorf("failed to estimate gas for CreatePid(): %s, address: %v", address, err)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// todo: remove authentication
	// call contract CreatePid()
	tx, err := s.didContractInstance.CreatePid(opts, createTime, "", publicKey, updateTime)
	if err != nil {
		log.WithError(err).Errorf("failed to call CreatePid(), address: %s", address)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.TxHash = tx.Hash()
	response.Status = Response_SUCCESS

	log.Debugf("call CreatePid() txHash:%s, addresss:%s", tx.Hash().Hex(), address)

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = Response_UNKNOWN
		response.Data = doc.BuildDid(addrHex)
		response.Msg = "failed to get tx receipt"
		return response
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Status = Response_FAILURE
		response.Msg = "failed to process tx"
	} else {
		response.Data = doc.BuildDid(addrHex)
	}

	return response
}

const (
	did_EVENT_FIELD_CREATE         uint8 = 0
	did_EVENT_FIELD_AUTHENTICATION uint8 = 1
	did_EVENT_FIELD_PUBLICKEY      uint8 = 2
)

func (s *DocumentService) GetDocument(address ethcommon.Address) *Response[*doc.DidDocument] {
	// init the result
	response := new(Response[*doc.DidDocument])
	response.CallMode = true
	response.Status = Response_SUCCESS

	blockNo, err := s.didContractInstance.GetLatestBlock(nil, address)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetLatestBlock(), address: %s", address)
		response.Status = Response_FAILURE
		response.Msg = "failed to get latest block of DID"
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("DID not found, address: %s", address)
		response.Status = Response_FAILURE
		response.Msg = "DID not found"
		return response
	}

	document := new(doc.DidDocument)
	document.Id = doc.BuildDid(address.Hex())

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	prevBlock := blockNo

	for prevBlock.Uint64() > 0 {
		logs := s.ctx.GetLog(timeoutCtx, didContractAddress, prevBlock)
		for _, eachLog := range logs {
			event, err := s.didContractInstance.ParsePIDAttributeChange(*eachLog)
			if err != nil {
				response.Status = Response_FAILURE
				response.Msg = "failed to parse contract event"
				return response
			}
			switch event.FieldKey {
			case did_EVENT_FIELD_CREATE:
				//NOP
				document.Created = event.FieldValue
				document.Updated = event.UpdateTime
				prevBlock = event.BlockNumber

			/*case did_EVENT_FIELD_AUTHENTICATION:
			//调用
			auths := make([]*doc.DidAuthentication, 0)
			if err := json.Unmarshal([]byte(event.FieldValue), &auths); err != nil {
				response.Status = common.Response_FAILURE
				response.Msg = "failed to unmarshal DIdAuthentication"
				return response
			}
			document.Updated = event.UpdateTime
			document.AddAuthentication(auths)

			prevBlock = event.BlockNumber*/
			case did_EVENT_FIELD_PUBLICKEY:
				pubKeys := make([]*doc.DidPublicKey, 0)
				if err := json.Unmarshal([]byte(event.FieldValue), &pubKeys); err != nil {
					response.Status = Response_FAILURE
					response.Msg = "failed to unmarshal DidPublicKey"
					return response
				}
				document.Updated = event.UpdateTime
				document.AddPublicKey(pubKeys)

				prevBlock = event.BlockNumber
			}
		}
	}
	response.Data = document

	return response

}

// 这是对合约setAttribute()方法的一个包装
func (s *DocumentService) AddPublicKey(address ethcommon.Address, pubKeyId string, keyType doc.PublicKeyType, publicKey string) *Response[bool] {
	response := new(Response[bool])
	response.CallMode = false
	response.Status = Response_SUCCESS

	response.Data = false

	//to check if DID document has ths public key id already?
	hasResp := s.HasPublicKey(address, pubKeyId)

	if hasResp.Status != Response_SUCCESS {
		response.Status = hasResp.Status
		response.Msg = hasResp.Msg
		return response
	} else if hasResp.Data == true {
		response.Status = Response_FAILURE
		response.Msg = "public key exists"
		return response
	}

	updateTime := common.FormatUTC(time.Now().UTC())

	fieldValue := doc.BuildDidPublicKeys(pubKeyId, publicKey, address.Hex(), keyType)
	input, err := PackAbiInput(s.abi, "setAttribute", did_EVENT_FIELD_PUBLICKEY, fieldValue, updateTime)

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
	tx, err := s.didContractInstance.SetAttribute(opts, did_EVENT_FIELD_PUBLICKEY, fieldValue, updateTime)
	if err != nil {
		log.WithError(err).Errorf("failed to call SetAttribute(), address: %s", address)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	response.TxHash = tx.Hash()
	response.Status = Response_SUCCESS

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
	} else {
		response.Data = true
	}
	return response
}

func (s *DocumentService) HasPublicKey(address ethcommon.Address, pubKeyId string) *Response[bool] {
	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_SUCCESS

	if docResp := s.GetDocument(address); docResp.Status != Response_SUCCESS {
		response.Msg = docResp.Msg
		response.Status = docResp.Status
		return response
	} else {
		doc := docResp.Data
		if doc.FindPublicKey(pubKeyId) != nil {
			response.Data = true
		} else {
			response.Data = false
		}
		return response
	}
}

func (s *DocumentService) GetPublicKey(address ethcommon.Address, pubKeyId string) *Response[*doc.DidPublicKey] {
	response := new(Response[*doc.DidPublicKey])
	response.CallMode = true
	response.Status = Response_SUCCESS

	if docResp := s.GetDocument(address); docResp.Status != Response_SUCCESS {
		response.Msg = docResp.Msg
		response.Status = docResp.Status
		return response
	} else {
		doc := docResp.Data
		if didPubKey := doc.FindPublicKey(pubKeyId); didPubKey != nil {
			response.Data = didPubKey
			response.Status = Response_SUCCESS
		} else {
			response.Status = Response_FAILURE
			return response
		}
		return response
	}
}
