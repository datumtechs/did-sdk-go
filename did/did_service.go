package did

import (
	"github.com/bglmmz/chainclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var (
	didContractAddress      = ethcommon.HexToAddress("0x2C9D6e8e8c1Ac6F171E49ea5b8C0b7e215D8f254")
	pctContractAddress      = ethcommon.HexToAddress("0x0101222470f5A7275F5C17aDFE89535DfF05c1DD")
	proposalContractAddress = ethcommon.HexToAddress("0x8Ae521006b24c1Bc586B824053350CEcf56AA12E")
	vcContractAddress       = ethcommon.HexToAddress("0xa3bbFdC63B1F78b1eC7C3Ce70E606b7173f666c6")
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
	SupplementDidPublicKey(address ethcommon.Address, pubKeyId string, keyType doc.PublicKeyType, PublicKey string) *Response[bool]
	HasPublicKey(address ethcommon.Address, pubKey string) *Response[bool]
	GetDidPublicKey(address ethcommon.Address, pubKey string) *Response[*doc.DidPublicKey]
}
func (s *DIDService) CreateDID(address ethcommon.Address, pubKeyHex string) *Response[string] {
	return s.documentService.CreateDID(address, pubKeyHex)
}*/
