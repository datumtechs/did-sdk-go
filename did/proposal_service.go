package did

import (
	"context"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/types/proposal"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

func (s *DIDService) GetAllAuthority(applicantDid string, applicant ethcommon.Address, pctId uint64, claim map[string]interface{}, issuer ethcommon.Address) *common.Response[[]proposal.Authority] {
	// init the result
	response := new(common.Response[[]proposal.Authority])
	response.CallMode = true

	addressList, urlList, err := s.proposalContractInstance.GetAllAuthority(nil)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetAllAuthority(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	if len(addressList) != len(urlList) {
		log.WithError(err).Errorf("data returned from GetAllAuthority() error")
		response.Status = common.Response_FAILURE
		response.Msg = "data returned from contract error"
		return response
	}

	authorityList := make([]proposal.Authority, len(addressList))
	for i := 0; i < len(addressList); i++ {
		authorityList[i].Address = addressList[i]
		authorityList[i].Url = urlList[i]
	}
	response.Data = authorityList
	return response
}

func (s *DIDService) SubmitProposal(proposalUrl string, proposed ethcommon.Address, rpcUrl string) *common.Response {
	// init the result
	response := new(common.Response)
	response.CallMode = false

	// prepare parameters for submitProposal()
	input, err := s.packInput("submitProposal", proposal.ProposalType_ADD, proposalUrl, proposed, rpcUrl)
	if err != nil {
		log.Errorf("failed to pack input data for submitProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, proposalContractAddress, input)
	if err != nil {
		log.Errorf("failed to estimate gas for submitProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.proposalContractInstance.SubmitProposal(opts, uint8(proposal.ProposalType_ADD), proposalUrl, proposed, rpcUrl)
	if err != nil {
		log.WithError(err).Errorf("failed to call submitProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.TxHash = tx.Hash()
	response.Status = common.Response_SUCCESS

	log.Debugf("call submitProposal() txHash: %s", tx.Hash().Hex())

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

		//todo: retrieve proposalId from log, and set to response.data
		for _, txLog := range receipt.Logs {
			if newProposalEvent, err := s.proposalContractInstance.ParseNewProposal(*txLog); err != nil {
				log.Debugf("newProposalEvent: %#v", newProposalEvent)
			}
		}
	}

	return response
}

func (s *DIDService) VoteProposal(proposalId *big.Int) *common.Response {
	// init the result
	response := new(common.Response)
	response.CallMode = false

	// prepare parameters for submitProposal()
	input, err := s.packInput("VoteProposal", proposalId)
	if err != nil {
		log.Errorf("failed to pack input data for VoteProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, proposalContractAddress, input)
	if err != nil {
		log.Errorf("failed to estimate gas for VoteProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.proposalContractInstance.VoteProposal(opts, proposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to call VoteProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.TxHash = tx.Hash()
	response.Status = common.Response_SUCCESS

	log.Debugf("call VoteProposal() txHash: %s", tx.Hash().Hex())

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

func (s *DIDService) WithdrawProposal(proposalId *big.Int) *common.Response {
	// init the result
	response := new(common.Response)
	response.CallMode = false

	// prepare parameters for submitProposal()
	input, err := s.packInput("WithdrawProposal", proposalId)
	if err != nil {
		log.Errorf("failed to pack input data for WithdrawProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}
	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, proposalContractAddress, input)
	if err != nil {
		log.Errorf("failed to estimate gas for WithdrawProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.proposalContractInstance.WithdrawProposal(opts, proposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to call WithdrawProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.TxHash = tx.Hash()
	response.Status = common.Response_SUCCESS

	log.Debugf("call WithdrawProposal() txHash: %s", tx.Hash().Hex())

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

func (s *DIDService) EffectProposal(proposalId *big.Int) *common.Response {
	// init the result
	response := new(common.Response)
	response.CallMode = false

	// prepare parameters for EffectProposal()
	input, err := s.packInput("EffectProposal", proposalId)
	if err != nil {
		log.Errorf("failed to pack input data for EffectProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to pack input data"
		return response
	}

	timeout := time.Duration(5000) * time.Millisecond
	timeoutCtx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	// 估算gas
	gasEstimated, err := s.ctx.EstimateGas(timeoutCtx, proposalContractAddress, input)
	if err != nil {
		log.Errorf("failed to estimate gas for EffectProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to estimate gas"
		return response
	}

	// 交易参数直接使用用户预付的总的gas，尽量放大，以防止交易执行gas不足
	gasEstimated = uint64(float64(gasEstimated) * 1.30)
	opts, err := s.ctx.BuildTxOpts(0, gasEstimated)

	// call contract CreatePid()
	tx, err := s.proposalContractInstance.EffectProposal(opts, proposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to call EffectProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.TxHash = tx.Hash()
	response.Status = common.Response_SUCCESS

	log.Debugf("call EffectProposal() txHash: %s", tx.Hash().Hex())

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

func (s *DIDService) GetAllProposalId() *common.Response[[]*big.Int] {
	// init the result
	response := new(common.Response[[]*big.Int])
	response.CallMode = true

	// call contract getAllProposalId()
	pIdList, err := s.proposalContractInstance.GetAllProposalId(nil)
	if err != nil {
		log.WithError(err).Errorf("failed to call getAllProposalId(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.Status = common.Response_SUCCESS
	response.Data = pIdList

	return response
}

func (s *DIDService) GetProposalId(blockNo uint64) *common.Response[[]*big.Int] {
	// init the result
	response := new(common.Response[[]*big.Int])
	response.CallMode = true

	// call contract getProposalId()
	pIdList, err := s.proposalContractInstance.GetProposalId(nil, new(big.Int).SetUint64(blockNo))
	if err != nil {
		log.WithError(err).Errorf("failed to call getProposalId(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}
	response.Status = common.Response_SUCCESS
	response.Data = pIdList

	return response
}

func (s *DIDService) GetProposal(proposalId *big.Int) *common.Response[*proposal.Proposal] {
	// init the result
	response := new(common.Response[*proposal.Proposal])
	response.CallMode = true

	// call contract getProposalId()
	pType, pUrl, candidate, candidateServiceUrl, submitter, submitBlockNo, err := s.proposalContractInstance.GetProposal(nil, proposalId)
	if err != nil {
		log.WithError(err).Errorf("failed to call GetProposal(), error: %+v", err)
		response.Status = common.Response_FAILURE
		response.Msg = "failed to call contract"
		return response
	}

	proposal := &proposal.Proposal{
		ProposalType:        pType,
		ProposalUrl:         pUrl,
		Submitter:           submitter,
		Candidate:           candidate,
		CandidateServiceUrl: candidateServiceUrl,
		SubmitBlockNo:       submitBlockNo.Uint64(),
	}

	response.Data = proposal
	response.Status = common.Response_SUCCESS
	return response
}
