package did

import (
	"github.com/datumtechs/did-sdk-go/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var candidateAddress = ethcommon.HexToAddress("0xEFb8aeE7c9BC8c8f1472976299855e7059b8Ecda")

func Test_GetAuthority(t *testing.T) {
	setup()
	response := didService.ProposalService.GetAuthority(address)
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)
}

func Test_GetAllAuthority(t *testing.T) {
	setup()
	response := didService.ProposalService.GetAllAuthority()
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)
}
func Test_SubmitProposal(t *testing.T) {
	setup()
	req := new(SubmitProposalReq)
	req.ProposalType = uint8(types.ProposalType_ADD)
	req.ProposalUrl = "http://proposal_add.url"
	req.Candidate = candidateAddress
	req.CandidateServiceUrl = "http://localhost:8458"
	req.PrivateKey = privateKey

	response := didService.ProposalService.SubmitProposal(*req)
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)
	// data: 0
	// txHash: 0x5f066270441dee02eedddfad7196d05e0d6759e66025c85f8c1525a46118b025
	// blockNumber: 26628186
}
func Test_GetAllProposalId(t *testing.T) {
	setup()
	response := didService.ProposalService.GetAllProposalId()
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)
}
func Test_GetProposalId(t *testing.T) {
	setup()

	response := didService.ProposalService.GetProposalId(26628186)
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)

}

func Test_GetProposal(t *testing.T) {
	setup()

	response := didService.ProposalService.GetProposal(big.NewInt(0))
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)

}

func Test_ResetInterval(t *testing.T) {
	setup()
	req := new(ResetIntervalReq)
	req.PrivateKey = privateKey

	req.IntervalType = types.ProposalIntervalType_Begin
	req.Blocks = big.NewInt(200)

	response := didService.ProposalService.ResetInterval(*req)
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)

}

func Test_VoteProposal(t *testing.T) {
	setup()
	req := new(VoteProposalReq)
	req.ProposalId = big.NewInt(0)
	req.PrivateKey = privateKey
	response := didService.ProposalService.VoteProposal(*req)
	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	t.Logf("resposne.data:%+v", response.Data)
}
