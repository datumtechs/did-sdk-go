package did

import (
	"github.com/bglmmz/chainclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var (
	didContractAddress      = ethcommon.HexToAddress("0xfd67957F61F9cC7A85da7657ED0B54b0A5867223")
	pctContractAddress      = ethcommon.HexToAddress("0xD4F004109ed8097cbb46dBD22334096DB37b5ce6")
	proposalContractAddress = ethcommon.HexToAddress("0x0a033cA8faA1cBB184d196a71A0BaeD76cE493cB")
	vcContractAddress       = ethcommon.HexToAddress("0xE9eD5A657eA3474FD2f3Cc40e93A3470A643E60E")
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
