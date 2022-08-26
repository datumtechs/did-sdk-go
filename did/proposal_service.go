package did

import (
	"context"
	"crypto/ecdsa"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/contracts"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
	"time"
)

type ProposalService struct {
	ctx                      chainclient.Context
	abi                      abi.ABI
	proposalContractInstance *contracts.Proposal
	proposalContractProxy    ethcommon.Address
}

func NewProposalService(ctx chainclient.Context, config *Config) *ProposalService {
	log.Info("Init Proposal Service ...")
	m := new(ProposalService)
	m.ctx = ctx
	m.proposalContractProxy = config.ProposalContractProxy

	instance, err := contracts.NewProposal(m.proposalContractProxy, ctx.GetClient())
	if err != nil {
		log.Fatal(err)
	}
	m.proposalContractInstance = instance

	abiCode, err := abi.JSON(strings.NewReader(contracts.ProposalMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	m.abi = abiCode
	return m
}

/*func (s *ProposalService) GetAuthority(bech32Addr ethcommon.Address) *Response[types.Authority] {
	// init the result
	response := new(Response[types.Authority])
	response.CallMode = true
	response.Status = Response_FAILURE

	bech32Addr, url, _, err := s.proposalContractInstance.GetAuthority(nil, bech32Addr)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetAuthority(), bech32Addr: %s", bech32Addr.String())
		response.Msg = "failed to call contract"
		return response
	}
	if bech32Addr == (ethcommon.Address{}) || len(url) > 0 {
		response.Status = Response_FAILURE
		response.Msg = "authority info is broken"
	}
	auth := types.Authority{}
	auth.Address = bech32Addr
	auth.Url = url
	response.Status = Response_SUCCESS
	response.Data = auth
	return response

}*/

func (s *ProposalService) GetAuthority(address ethcommon.Address) *Response[types.Authority] {
	// init the result
	response := new(Response[types.Authority])
	response.CallMode = true
	response.Status = Response_FAILURE

	resp := s.GetAllAuthority()
	if resp.Status != Response_SUCCESS {
		CopyResp(resp, response)
		return response
	}
	for _, item := range resp.Data {
		if item.Address == address {
			response.Data = item
			response.Status = Response_SUCCESS
			return response
		}
	}
	response.Status = Response_FAILURE
	response.Msg = "Did not found"
	return response

}

func (s *ProposalService) GetAllAuthority() *Response[[]types.Authority] {
	// init the result
	response := new(Response[[]types.Authority])
	response.CallMode = true
	response.Status = Response_FAILURE

	addressList, urlList, _, err := s.proposalContractInstance.GetAllAuthority(nil)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetAllAuthority(), error: %+v", err)

		response.Msg = "failed to call contract"
		return response
	}

	if len(addressList) != len(urlList) {
		log.WithError(err).Errorf("data returned from GetAllAuthority() error")
		response.Msg = "data returned from contract error"
		return response
	}

	authorityList := make([]types.Authority, len(addressList))
	for i := 0; i < len(addressList); i++ {
		authorityList[i].Address = addressList[i]
		authorityList[i].Url = urlList[i]
	}
	response.Data = authorityList
	response.Status = Response_SUCCESS
	return response
}

type SubmitProposalReq struct {
	PrivateKey          *ecdsa.PrivateKey `json:"-"` // Required: The private key to sign transaction
	ProposalType        uint8
	ProposalUrl         string
	Candidate           ethcommon.Address
	CandidateServiceUrl string
}

func (s *ProposalService) SubmitProposal(req SubmitProposalReq) *Response[string] {
	// init the result
	response := new(Response[string])
	response.CallMode = false
	response.Status = Response_FAILURE
	response.Data = "0"

	// init the tx signer
	s.ctx.SetPrivateKey(req.PrivateKey)

	// prepare parameters for submitProposal()
	input, err := PackAbiInput(s.abi, "submitProposal", req.ProposalType, req.ProposalUrl, req.Candidate, req.CandidateServiceUrl)
	if err != nil {
		log.Errorf("failed to pack input data for submitProposal(), error: %+v", err)
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.proposalContractProxy, input)
	if err != nil {
		log.Errorf("failed to estimate gas for submitProposal(), error: %+v", err)
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract SubmitProposal()
	tx, err := s.proposalContractInstance.SubmitProposal(opts, req.ProposalType, req.ProposalUrl, req.Candidate, req.CandidateServiceUrl)
	if err != nil {
		log.WithError(err).Errorf("failed to call submitProposal(), error: %+v", err)
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("call submitProposal() txHash: %s", tx.Hash().Hex())

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
			if newProposalEvent, err := s.proposalContractInstance.ParseNewProposal(*txLog); err == nil {
				log.Debugf("newProposalEvent: %#v", newProposalEvent)
				if newProposalEvent.Candidate == req.Candidate {
					response.Data = newProposalEvent.ProposalId.String()
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

type VoteProposalReq struct {
	// Required: The private key to sign transaction
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	ProposalId *big.Int
}

func (s *ProposalService) VoteProposal(req VoteProposalReq) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = false
	response.Status = Response_FAILURE
	response.Data = false

	// init the tx signer
	s.ctx.SetPrivateKey(req.PrivateKey)

	// prepare parameters for submitProposal()
	input, err := PackAbiInput(s.abi, "voteProposal", req.ProposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to pack input data for VoteProposal(),proposalId:%d", req.ProposalId)
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.proposalContractProxy, input)
	if err != nil {
		log.WithError(err).Errorf("failed to estimate gas for VoteProposal(),proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract VoteProposal()
	tx, err := s.proposalContractInstance.VoteProposal(opts, req.ProposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to call VoteProposal(),proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("call VoteProposal() proposalId:%d txHash: %s", req.ProposalId, tx.Hash().Hex())
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
	response.Data = true
	response.Status = Response_SUCCESS
	return response
}

type WithdrawProposalReq struct {
	// Required: The private key to sign transaction
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	ProposalId *big.Int
}

func (s *ProposalService) WithdrawProposal(req WithdrawProposalReq) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = false
	response.Status = Response_FAILURE
	response.Data = false

	// init the tx signer
	s.ctx.SetPrivateKey(req.PrivateKey)

	// prepare parameters for WithdrawProposal()
	input, err := PackAbiInput(s.abi, "withdrawProposal", req.ProposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to pack input data for WithdrawProposal(),proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}
	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.proposalContractProxy, input)
	if err != nil {
		log.WithError(err).Errorf("failed toto estimate gas for WithdrawProposal(),proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract WithdrawProposal()
	tx, err := s.proposalContractInstance.WithdrawProposal(opts, req.ProposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to call WithdrawProposal(),proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("call WithdrawProposal() proposalId:%d txHash: %s", req.ProposalId, tx.Hash().Hex())

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
	response.Data = true
	response.Status = Response_SUCCESS
	return response
}

type EffectProposalReq struct {
	// Required: The private key to sign transaction
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	ProposalId *big.Int
}

func (s *ProposalService) EffectProposal(req EffectProposalReq) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = false
	response.Status = Response_FAILURE
	response.Data = false

	// init the tx signer
	s.ctx.SetPrivateKey(req.PrivateKey)

	// prepare parameters for EffectProposal()
	input, err := PackAbiInput(s.abi, "effectProposal", req.ProposalId)
	if err != nil {
		log.WithError(err).Errorf("EffectProposal: failed to pack input data,proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.proposalContractProxy, input)
	if err != nil {
		log.WithError(err).Errorf("EffectProposal: failed to estimate gas,proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract EffectProposal()
	tx, err := s.proposalContractInstance.EffectProposal(opts, req.ProposalId)
	if err != nil {
		log.WithError(err).Errorf("EffectProposal: failed to call contract, proposalId:%d", req.ProposalId)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("CreateEvidence: call contract txHash: %s", tx.Hash().Hex())

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
	response.Data = true
	response.Status = Response_SUCCESS
	return response
}

func (s *ProposalService) GetAllProposalId() *Response[[]*big.Int] {
	// init the result
	response := new(Response[[]*big.Int])
	response.CallMode = true
	response.Status = Response_FAILURE

	// call contract getAllProposalId()
	pIdList, err := s.proposalContractInstance.GetAllProposalId(nil)
	if err != nil {
		log.WithError(err).Error("failed to call getAllProposalId()")
		response.Msg = "failed to call contract"
		return response
	}
	response.Status = Response_SUCCESS
	response.Data = pIdList
	return response
}

func (s *ProposalService) GetProposalId(blockNo uint64) *Response[[]*big.Int] {
	// init the result
	response := new(Response[[]*big.Int])
	response.CallMode = true
	response.Status = Response_FAILURE

	// call contract getProposalId()
	pIdList, err := s.proposalContractInstance.GetProposalId(nil, new(big.Int).SetUint64(blockNo))
	if err != nil {
		log.WithError(err).Error("failed to call getProposalId()")
		response.Msg = "failed to call contract"
		return response
	}
	response.Status = Response_SUCCESS
	response.Data = pIdList
	return response
}

func (s *ProposalService) GetProposal(proposalId *big.Int) *Response[*types.Proposal] {
	// init the result
	response := new(Response[*types.Proposal])
	response.CallMode = true
	response.Status = Response_FAILURE

	// call contract getProposalId()
	pType, pUrl, candidate, candidateServiceUrl, submitter, submitBlockNo, _, err := s.proposalContractInstance.GetProposal(nil, proposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetProposal(),proposalId:%d", proposalId)
		response.Msg = "failed to call contract"
		return response
	}

	proposal := &types.Proposal{
		ProposalType:        pType,
		ProposalUrl:         pUrl,
		Submitter:           submitter,
		Candidate:           candidate,
		CandidateServiceUrl: candidateServiceUrl,
		SubmitBlockNo:       submitBlockNo.Uint64(),
	}
	response.Status = Response_SUCCESS
	response.Data = proposal
	return response
}

type ResetIntervalReq struct {
	// Required: The private key to sign transaction
	PrivateKey   *ecdsa.PrivateKey
	IntervalType types.ProposalIntervalType
	Blocks       *big.Int
}

func (s *ProposalService) ResetInterval(req ResetIntervalReq) *Response[bool] {
	// init the result
	response := new(Response[bool])
	response.CallMode = false
	response.Status = Response_FAILURE
	response.Data = false

	// init the tx signer
	s.ctx.SetPrivateKey(req.PrivateKey)

	// prepare parameters for EffectProposal()
	input, err := PackAbiInput(s.abi, "setInterval", uint8(req.IntervalType), req.Blocks)
	if err != nil {
		log.WithError(err).Errorf("ResetInterval: failed to pack input data,IntervalType:%d", req.IntervalType)
		response.Status = Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(10000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, s.proposalContractProxy, input)
	if err != nil {
		log.WithError(err).Errorf("ResetInterval: failed to estimate gas,IntervalType:%d", req.IntervalType)
		response.Status = Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract EffectProposal()
	tx, err := s.proposalContractInstance.SetInterval(opts, uint8(req.IntervalType), req.Blocks)
	if err != nil {
		log.WithError(err).Errorf("ResetInterval: failed to call contract,IntervalType:%d", req.IntervalType)
		response.Status = Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	log.Debugf("ResetInterval: call contract txHash: %s", tx.Hash().Hex())

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
	response.Data = true
	response.Status = Response_SUCCESS
	return response

}
