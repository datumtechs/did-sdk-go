package did

import (
	"crypto/ecdsa"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
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
	PrivateKey *ecdsa.PrivateKey // the private key to sign the credential
	PctJson    string
	Extra      []byte
}

func (s *PctService) GetPct(pctId *big.Int) *Response[*types.Pct] {
	// init the result
	response := new(Response[*types.Pct])
	response.CallMode = true
	response.Status = Response_SUCCESS

	issuer, jsonSchema, extra, err := s.pctContractInstance.GetPctInfo(nil, pctId)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetPctInfo(), pctId: %d", pctId)
		response.Status = Response_FAILURE
		response.Msg = "failed to get PCT"
		return response
	} else if len(jsonSchema) == 0 {
		log.WithError(err).Errorf("cannot find pct, pctId: %d", pctId)
		response.Status = Response_FAILURE
		response.Msg = "cannot find PCT"
		return response
	}

	pctObj := new(types.Pct)
	pctObj.Issuer = issuer
	pctObj.JsonSchema = jsonSchema
	pctObj.Extra = extra
	response.Data = pctObj
	return response
}

func (s *PctService) VerifyByPct(pctId *big.Int, content map[string]interface{}) *Response[bool] {

	response := new(Response[bool])
	response.CallMode = true
	response.Status = Response_SUCCESS

	response.Data = false

	pctResp := s.GetPct(pctId)
	if pctResp.Status != Response_SUCCESS {
		CopyResp(pctResp, response)
		return response
	}
	pctObj := pctResp.Data
	jsonSchema := pctObj.JsonSchema

	response.Data = common.VerifyWithJsonSchema(jsonSchema, content)

	return response
}
