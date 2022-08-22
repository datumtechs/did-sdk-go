package did

import (
	"context"
	"crypto/ecdsa"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
	"time"
)

type PctService struct {
	ctx                 chainclient.Context
	abi                 abi.ABI
	pctContractInstance *contracts.Pct
}

func NewPctService(ctx chainclient.Context) *PctService {
	log.Info("Init Pct Service ...")
	m := new(PctService)
	m.ctx = ctx

	instance, err := contracts.NewPct(pctContractAddress, ctx.GetClient())
	if err != nil {
		log.Fatal(err)
	}
	m.pctContractInstance = instance

	abiCode, err := abi.JSON(strings.NewReader(contracts.PctMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	m.abi = abiCode
	return m
}

type CreatePctReq struct {
	PrivateKey *ecdsa.PrivateKey `json:"-"` // the private key to sign the credential
	PctJson    string
	Extra      []byte
}

func (s *PctService) RegisterPct(req CreatePctReq) *Response[string] { // init the result
	response := new(Response[string])
	response.CallMode = false
	response.Status = Response_FAILURE

	// init the tx signer
	s.ctx.SetPrivateKey(req.PrivateKey)

	if !common.VerifyJsonSchema(req.PctJson) {
		response.Msg = "pct json invalid"
		return response
	}

	input, err := PackAbiInput(s.abi, "registerPct", req.PctJson, req.Extra)
	if err != nil {
		log.WithError(err).Errorf("failed topack input data for registerPct(), pct:%s", req.PctJson)
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, pctContractAddress, input)
	if err != nil {
		log.WithError(err).Errorf("failed to estimate gas for registerPct(), pct:%s", req.PctJson)
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)
	if err != nil {
		log.WithError(err).Errorf("failed to estimate gas for registerPct(), pct:%s", req.PctJson)
		response.Msg = "failed to estimate gas"
		return response
	}

	// call contract CreateDid()
	tx, err := s.pctContractInstance.RegisterPct(opts, req.PctJson, req.Extra)
	if err != nil {
		log.WithError(err).Errorf("failed to call registerPct(), pct:%s", req.PctJson)
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("call registerPct() txHash:%s, pct:%s", tx.Hash().Hex(), req.PctJson)

	// to get receipt and assemble result
	receipt := s.ctx.WaitReceipt(timeoutCtx, tx.Hash(), time.Duration(500)*time.Millisecond) // period 500 ms
	if nil == receipt {
		response.Status = Response_UNKNOWN
		response.Msg = "failed to get tx receipt"
		return response
	}

	// contract tx execute failed.
	if receipt.Status == 0 {
		response.Msg = "failed to process tx"
		return response
	} else {
		// retrieve proposalId from log, and set to response.data
		for _, txLog := range receipt.Logs {
			if registerPctEvent, err := s.pctContractInstance.ParseRegisterPct(*txLog); err == nil {
				log.Debugf("registerPctEvent: %#v", registerPctEvent)
				if registerPctEvent.JsonSchema == req.PctJson {
					response.Data = registerPctEvent.PctId.String()
					//shortcut the circle
					break
				}
			}
		}
	}

	// 交易信息
	response.TxInfo = NewTransactionInfo(receipt)
	response.Status = Response_SUCCESS
	return response
}

func (s *PctService) GetPct(pctId *big.Int) *Response[*types.Pct] {
	// init the result
	response := new(Response[*types.Pct])
	response.CallMode = true
	response.Status = Response_FAILURE

	issuer, jsonSchema, extra, err := s.pctContractInstance.GetPctInfo(nil, pctId)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetPctInfo(), pctId: %d", pctId)
		response.Msg = "failed to get PCT"
		return response
	} else if len(jsonSchema) == 0 {
		log.WithError(err).Errorf("cannot find pct, pctId: %d", pctId)
		response.Msg = "cannot find PCT"
		return response
	}

	pctObj := new(types.Pct)
	pctObj.Issuer = issuer
	pctObj.JsonSchema = jsonSchema
	pctObj.Extra = extra

	response.Status = Response_SUCCESS
	response.Data = pctObj

	return response
}

func (s *PctService) VerifyByPct(pctId *big.Int, content map[string]interface{}) *Response[bool] {

	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_FAILURE
	response.Data = false

	pctResp := s.GetPct(pctId)
	if pctResp.Status != Response_SUCCESS {
		CopyResp(pctResp, response)
		return response
	}
	pctObj := pctResp.Data
	jsonSchema := pctObj.JsonSchema

	response.Data = common.VerifyWithJsonSchema(jsonSchema, content)
	response.Status = Response_SUCCESS
	return response
}
