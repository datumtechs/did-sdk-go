package did

import (
	"github.com/datumtechs/chainclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

/*var (
	documentContractProxy   = ethcommon.HexToAddress("0x279167d9767b10CEF88b9a81D9C912e475c0B75b")
	pctContractProxy        = ethcommon.HexToAddress("0xFa2A71584740c749F1EF055741140833534504CD")
	proposalContractProxy   = ethcommon.HexToAddress("0x857027b23F73F5823984d90550A39cAA6FA43A11")
	credentialContractProxy = ethcommon.HexToAddress("0x6Afe474d2201525558b50D5D5a8544e88d222E05")
)
*/
type Config struct {
	DocumentContractProxy   ethcommon.Address
	PctContractProxy        ethcommon.Address
	ProposalContractProxy   ethcommon.Address
	CredentialContractProxy ethcommon.Address
}

func PackAbiInput(abi abi.ABI, method string, params ...interface{}) ([]byte, error) {
	return abi.Pack(method, params...)
}

type DIDService struct {
	Config            *Config
	DocumentService   *DocumentService
	PctService        *PctService
	ProposalService   *ProposalService
	CredentialService *CredentialService
}

func NewDIDService(ctx chainclient.Context, config *Config) *DIDService {
	log.Info("Init DID service ...")

	didService := new(DIDService)

	didService.Config = config

	didService.DocumentService = NewDocumentService(ctx, config)
	didService.PctService = NewPctService(ctx, config)
	didService.ProposalService = NewProposalService(ctx, config)
	didService.CredentialService = NewCredentialService(ctx, config, didService.DocumentService, didService.PctService)
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
