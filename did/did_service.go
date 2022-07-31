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
	DocumentService *DocumentService
	PctService      *PctService
	ProposalService *ProposalService
	VcService       *VcService
}

func NewDIDService(ctx chainclient.Context) *DIDService {
	log.Info("Init DID service ...")

	didService := new(DIDService)
	didService.DocumentService = NewDocumentService(ctx)
	didService.PctService = NewPctService(ctx)
	didService.ProposalService = NewProposalService(ctx)
	didService.VcService = NewVcService(ctx, didService.DocumentService, didService.PctService)
	return didService
}

/*type DocuementService interface {
	CreateDID(address ethcommon.Address, pubKeyHex string) *Response[string]
	GetDocument(address ethcommon.Address) *Response[*doc.DidDocument]
	AddDidPublicKey(address ethcommon.Address, pubKeyId string, keyType doc.PublicKeyType, PublicKey string) *Response[bool]
	HasPublicKey(address ethcommon.Address, pubKey string) *Response[bool]
	GetDidPublicKey(address ethcommon.Address, pubKey string) *Response[*doc.DidPublicKey]
}
func (s *DIDService) CreateDID(address ethcommon.Address, pubKeyHex string) *Response[string] {
	return s.documentService.CreateDID(address, pubKeyHex)
}*/
