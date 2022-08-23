package did

import (
	"github.com/bglmmz/chainclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var (
	didContractAddress        = ethcommon.HexToAddress("0x279167d9767b10CEF88b9a81D9C912e475c0B75b")
	pctContractAddress        = ethcommon.HexToAddress("0xFa2A71584740c749F1EF055741140833534504CD")
	proposalContractAddress   = ethcommon.HexToAddress("0x857027b23F73F5823984d90550A39cAA6FA43A11")
	credentialContractAddress = ethcommon.HexToAddress("0x6Afe474d2201525558b50D5D5a8544e88d222E05")
)

func PackAbiInput(abi abi.ABI, method string, params ...interface{}) ([]byte, error) {
	return abi.Pack(method, params...)
}

type DIDService struct {
	DocumentService   *DocumentService
	PctService        *PctService
	ProposalService   *ProposalService
	CredentialService *CredentialService
}

func NewDIDService(ctx chainclient.Context) *DIDService {
	log.Info("Init DID service ...")

	didService := new(DIDService)
	didService.DocumentService = NewDocumentService(ctx)
	didService.PctService = NewPctService(ctx)
	didService.ProposalService = NewProposalService(ctx)
	didService.CredentialService = NewCredentialService(ctx, didService.DocumentService, didService.PctService)
	return didService
}

/*type DocuementService interface {
	CreateDID(bech32Addr ethcommon.Address, pubKeyHex string) *Response[string]
	GetDocument(bech32Addr ethcommon.Address) *Response[*doc.DidDocument]
	SupplementDidPublicKey(bech32Addr ethcommon.Address, pubKeyId string, keyType doc.PublicKeyType, PublicKey string) *Response[bool]
	HasPublicKey(bech32Addr ethcommon.Address, pubKey string) *Response[bool]
	GetDidPublicKey(bech32Addr ethcommon.Address, pubKey string) *Response[*doc.DidPublicKey]
}
func (s *DIDService) CreateDID(bech32Addr ethcommon.Address, pubKeyHex string) *Response[string] {
	return s.documentService.CreateDID(bech32Addr, pubKeyHex)
}*/
