package did

import (
	"context"
	"encoding/json"
	"github.com/bglmmz/chainclient"
	common "github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/types/doc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var (
	didContractAddress      = ethcommon.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
	proposalContractAddress = ethcommon.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
	vcContractAddress       = ethcommon.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
)

type DIDService struct {
	ctx                      chainclient.Context
	abi                      abi.ABI
	didContractInstance      *contracts.Did
	proposalContractInstance *contracts.Vote
	vcContractInstance       *contracts.Credential
}

func NewDIDService(ctx chainclient.Context) *DIDService {
	log.Info("Init DID service ...")
	m := new(DIDService)
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

func (s *DIDService) packInput(method string, params ...interface{}) ([]byte, error) {
	return s.abi.Pack(method, params...)
}

// datum 项目中，address就是carrier内置钱包地址。
func (s *DIDService) CreateDID(address ethcommon.Address, pubKeyHex string) *common.Response[string] {
	// init the result
	response := new(common.Response[string])
	response.CallMode = false

	// prepare parameters for CreatePid()
	addrHex := address.Hex()
	now := time.Now().UTC()
	createTime := common.FormatUTC(now)
	updateTime := createTime

	pubKeyId := doc.BuildPublicKeyId(addrHex, 1)
	publicKey := doc.BuildDidPublicKeys(pubKeyId, pubKeyHex, addrHex, doc.PublicKey_SECP256K1)

	input, err := s.packInput("createPid", createTime, publicKey, updateTime)
	if err != nil {
		log.Errorf("failed to pack input data for CreatePid(): %+v", err)
		response.Status = common.Response_FAILURE
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
		response.Status = common.Response_FAILURE
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
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.TxHash = tx.Hash()
	response.Status = common.Response_SUCCESS

	log.Debugf("call CreatePid() txHash:%s, addresss:%s", tx.Hash().Hex(), address)

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = common.Response_UNKNOWN
		response.Data = doc.BuildDid(addrHex)
		response.Msg = "failed to get tx receipt"
		return response
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Status = common.Response_FAILURE
		response.Msg = "failed to process tx"
	} else {
		response.Status = common.Response_SUCCESS
		response.Data = doc.BuildDid(addrHex)
	}

	return response
}

const (
	did_EVENT_FIELD_CREATE         uint8 = 0
	did_EVENT_FIELD_AUTHENTICATION uint8 = 1
	did_EVENT_FIELD_PUBLICKEY      uint8 = 2
)

func (s *DIDService) GetDocument(address ethcommon.Address) *common.Response[*doc.DidDocument] {
	// init the result
	response := new(common.Response[*doc.DidDocument])
	response.CallMode = true

	blockNo, err := s.didContractInstance.GetLatestBlock(nil, address)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetLatestBlock(), address: %s", address)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to get latest block of DID"
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("DID not found, address: %s", address)
		response.Status = common.Response_FAILURE
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
				response.Status = common.Response_FAILURE
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
					response.Status = common.Response_FAILURE
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
func (s *DIDService) AddPublicKey(address ethcommon.Address, pubKeyId string, keyType doc.PublicKeyType, publicKey string) *common.Response {
	response := new(common.Response)
	response.CallMode = false

	//to check if DID document has ths public key id already?
	hasResp := s.HasPublicKey(address, pubKeyId)

	if hasResp.Status != common.Response_SUCCESS {
		response.Status = hasResp.Status
		response.Msg = hasResp.Msg
		return response
	} else if hasResp.Data == true {
		response.Status = common.Response_FAILURE
		response.Msg = "public key exists"
		return response
	}

	updateTime := common.FormatUTC(time.Now().UTC())

	fieldValue := doc.BuildDidPublicKeys(pubKeyId, publicKey, address.Hex(), keyType)
	input, err := s.packInput("setAttribute", did_EVENT_FIELD_PUBLICKEY, fieldValue, updateTime)

	if err != nil {
		log.Errorf("failed to pack input data for SetAttribute(): %+v", err)
		response.Status = common.Response_FAILURE
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
		response.Status = common.Response_FAILURE
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
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	response.TxHash = tx.Hash()
	response.Status = common.Response_SUCCESS

	log.Debugf("call SetAttribute() txHash:%s, addresss:%s", tx.Hash().Hex(), address)

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = common.Response_UNKNOWN
		response.Msg = "failed to get tx receipt"
		return response
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Status = common.Response_FAILURE
		response.Msg = "failed to process tx"
	} else {
		response.Status = common.Response_SUCCESS
	}

	return response
}

func (s *DIDService) HasPublicKey(address ethcommon.Address, pubKeyId string) *common.Response[bool] {
	response := new(common.Response[bool])
	response.CallMode = true

	if docResp := s.GetDocument(address); docResp.Status != common.Response_SUCCESS {
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

func (s *DIDService) GetPublicKey(address ethcommon.Address, pubKeyId string) *common.Response[*doc.DidPublicKey] {
	response := new(common.Response[*doc.DidPublicKey])
	response.CallMode = true

	if docResp := s.GetDocument(address); docResp.Status != common.Response_SUCCESS {
		response.Msg = docResp.Msg
		response.Status = docResp.Status
		return response
	} else {
		doc := docResp.Data
		if didPubKey := doc.FindPublicKey(pubKeyId); didPubKey != nil {
			response.Data = didPubKey
			response.Status = common.Response_SUCCESS
		} else {
			response.Status = common.Response_FAILURE
			return response
		}
		return response
	}
}
