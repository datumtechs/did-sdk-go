package did

import (
	"github.com/bglmmz/chainclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var (
	didContractAddress      = ethcommon.HexToAddress("0xaF90C0e740B2741c227E6d11daF001CfE819ec31")
	pctContractAddress      = ethcommon.HexToAddress("0xb99De0f8d5AebDE6bF1EbCfc8E9Ce1a0f5b61A47")
	proposalContractAddress = ethcommon.HexToAddress("0x47cbAc8e7bA083459d9D641505c7b9695d9476C9")
	vcContractAddress       = ethcommon.HexToAddress("0x2A38BEb5760395980d3211E301bDf8EBAA3E36E4")
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
