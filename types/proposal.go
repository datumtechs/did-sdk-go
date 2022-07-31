package types

import "github.com/ethereum/go-ethereum/common"

type ProposalType uint8

const (
	ProposalType_ADD     ProposalType = 1
	ProposalType_KICKOUT ProposalType = 2
	ProposalType_QUIT    ProposalType = 3
)

type Proposal struct {
	ProposalType        uint8
	ProposalUrl         string
	Submitter           common.Address
	Candidate           common.Address
	CandidateServiceUrl string
	SubmitBlockNo       uint64
}

type Authority struct {
	Address common.Address
	Url     string
}
