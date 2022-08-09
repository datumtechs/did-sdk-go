package types

import ethcommon "github.com/ethereum/go-ethereum/common"

type ProposalType uint8

const (
	ProposalType_ADD     ProposalType = 1
	ProposalType_KICKOUT ProposalType = 2
	ProposalType_QUIT    ProposalType = 3
)

type ProposalIntervalType uint8

const (
	ProposalIntervalType_Begin    ProposalIntervalType = 1
	ProposalIntervalType_Duration ProposalIntervalType = 2
	ProposalIntervalType_Quit     ProposalIntervalType = 4
)

/*func (s ProposalIntervalType) Uint8() uint8 {
	switch s {
	case ProposalIntervalType_Begin:
		return 1
	case ProposalIntervalType_Duration:
		return 2
	case ProposalIntervalType_Quit:
		return 4
	default:
		return 0
	}
}*/

type Proposal struct {
	ProposalType        uint8
	ProposalUrl         string
	Submitter           ethcommon.Address
	Candidate           ethcommon.Address
	CandidateServiceUrl string
	SubmitBlockNo       uint64
}

type Authority struct {
	Address ethcommon.Address
	Url     string
}
