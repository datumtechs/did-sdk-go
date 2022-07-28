package did

import (
	"github.com/bglmmz/chainclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var (
	didContractAddress      = ethcommon.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
	pctContractAddress      = ethcommon.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
	proposalContractAddress = ethcommon.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
	vcContractAddress       = ethcommon.HexToAddress("0x263B1D39843BF2e1DA27d827e749992fbD1f1577")
)

func PackAbiInput(abi abi.ABI, method string, params ...interface{}) ([]byte, error) {
	return abi.Pack(method, params...)
}

type DIDService struct {
	documentService *DocumentService
	pctService      *PctService
	proposalService *ProposalService
	vcService       *VcService
}

func NewDIDService(ctx chainclient.Context) *DIDService {
	log.Info("Init DID service ...")

	didService := new(DIDService)
	didService.documentService = NewDocumentService(ctx)
	didService.pctService = NewPctService(ctx)
	didService.proposalService = NewProposalService(ctx)
	didService.vcService = NewVcService(ctx, didService.documentService, didService.pctService)
	return didService
}
