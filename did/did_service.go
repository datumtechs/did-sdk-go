package did

import (
	"context"
	"encoding/json"
	"github.com/bglmmz/chainclient"
	common2 "github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var (
	didContractAddress      = common.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
	proposalContractAddress = common.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
	vcContractAddress       = common.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
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
func (s *DIDService) CreateDID(address common.Address, pubKeyHex string) *common2.Response {
	// init the result
	response := new(common2.Response)
	response.CallMode = false

	// prepare parameters for CreatePid()
	addrHex := address.Hex()
	now := time.Now().UTC()
	createTime := common2.FormatUTC(now)
	updateTime := createTime

	authentication := types.BuildDidAuthentications(pubKeyHex, addrHex, types.Authentication_VALID)
	pubKeyId := types.BuildPublicKeyId(addrHex, 1)
	publicKey := types.BuildDidPublicKeys(pubKeyId, pubKeyHex, addrHex, types.PublicKey_SECP256K1, types.PublicKey_VALID)

	input, err := s.packInput("createPid", createTime, authentication, publicKey, updateTime)
	if err != nil {
		log.Errorf("failed to pack input data for CreatePid(): %+v", err)
		response.Status = common2.Response_FAILURE
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
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.didContractInstance.CreatePid(opts, createTime, authentication, publicKey, updateTime)
	if err != nil {
		log.WithError(err).Errorf("failed to call CreatePid(), address: %s", address)
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.TxHash = tx.Hash()
	response.Status = common2.Response_SUCCESS

	log.Debugf("call CreatePid() txHash:%s, addresss:%s", tx.Hash().Hex(), address)

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = common2.Response_UNKNOWN
		response.Data = types.BuildPid(addrHex)
		response.Msg = "failed to get tx receipt"
		return response
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to process tx"
	} else {
		response.Status = common2.Response_SUCCESS
		response.Data = types.BuildPid(addrHex)
	}

	return response
}

const (
	did_EVENT_FIELD_CREATE         uint8 = 0
	did_EVENT_FIELD_AUTHENTICATION uint8 = 1
	did_EVENT_FIELD_PUBLICKEY      uint8 = 2
)

func (s *DIDService) GetDocument(address common.Address) *common2.Response {
	// init the result
	response := new(common2.Response)
	response.CallMode = true

	blockNo, err := s.didContractInstance.GetLatestBlock(nil, address)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetLatestBlock(), address: %s", address)
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to get latest block of DID"
		return response
	}
	if blockNo == nil || blockNo.Uint64() == 0 {
		log.WithError(err).Errorf("DID not found, address: %s", address)
		response.Status = common2.Response_FAILURE
		response.Msg = "DID not found"
		return response
	}

	document := new(types.DidDocument)
	document.Id = types.BuildPid(address.Hex())

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	prevBlock := blockNo

	for prevBlock.Uint64() > 0 {
		logs := s.ctx.GetLog(timeoutCtx, didContractAddress, prevBlock)
		for _, eachLog := range logs {
			event, err := s.didContractInstance.ParsePIDAttributeChange(*eachLog)
			if err != nil {
				response.Status = common2.Response_FAILURE
				response.Msg = "failed to parse contract event"
				return response
			}
			switch event.FieldKey {
			case did_EVENT_FIELD_CREATE:
				//NOP
				document.Create = event.FieldValue
				document.Updated = event.UpdateTime
				prevBlock = event.BlockNumber

			case did_EVENT_FIELD_AUTHENTICATION:
				//调用
				auths := make([]types.DidAuthentication, 0)
				if err := json.Unmarshal([]byte(event.FieldValue), &auths); err != nil {
					response.Status = common2.Response_FAILURE
					response.Msg = "failed to unmarshal DIdAuthentication"
					return response
				}
				document.Updated = event.UpdateTime
				document.AddAuthentication(auths)

				prevBlock = event.BlockNumber
			case did_EVENT_FIELD_PUBLICKEY:
				pubKeys := make([]types.DidPublicKey, 0)
				if err := json.Unmarshal([]byte(event.FieldValue), &pubKeys); err != nil {
					response.Status = common2.Response_FAILURE
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
func (s *DIDService) AddPublicKey(address common.Address, pubKeyId string, keyType types.PublicKeyType, publicKey string) *common2.Response {
	//to check if DID document has ths public key id already?
	response := s.HasPublicKey(address, pubKeyId)

	if response.Status != common2.Response_SUCCESS {
		return response
	} else if has, ok := response.Data.(bool); ok && has {
		response.Status = common2.Response_FAILURE
		response.Msg = "public key exists"
		return response
	}

	updateTime := common2.FormatUTC(time.Now().UTC())

	fieldValue := types.BuildDidPublicKeys(pubKeyId, publicKey, address.Hex(), keyType, types.PublicKey_VALID)
	input, err := s.packInput("setAttribute", did_EVENT_FIELD_PUBLICKEY, fieldValue, updateTime)

	if err != nil {
		log.Errorf("failed to pack input data for SetAttribute(): %+v", err)
		response.Status = common2.Response_FAILURE
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
		response.Status = common2.Response_FAILURE
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
		response.Status = common2.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	response.TxHash = tx.Hash()
	response.Status = common2.Response_SUCCESS

	log.Debugf("call SetAttribute() txHash:%s, addresss:%s", tx.Hash().Hex(), address)

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = common2.Response_UNKNOWN
		response.Msg = "failed to get tx receipt"
		return response
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

func (s *DIDService) HasPublicKey(address common.Address, pubKeyId string) *common2.Response {
	response := s.GetDocument(address)
	if response.Status != common2.Response_SUCCESS {
		return response
	}

	doc, ok := response.Data.(types.DidDocument)
	if !ok {
		response.Msg = "cannot cast data to DidDocument"
		response.Status = common2.Response_FAILURE
		return response
	}

	response.Status = common2.Response_SUCCESS
	if doc.HasPublicKey(pubKeyId) {
		response.Data = true
	} else {
		response.Data = false
	}
	return response
}
