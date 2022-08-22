package did

import (
	"github.com/bglmmz/chainclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var (
	didContractAddress        = ethcommon.HexToAddress("0x043Ed3C631c6d27fe83093E50Eb60b173939F8B3")
	pctContractAddress        = ethcommon.HexToAddress("0xCddB173756073A6C110886a4558cA4e8841fb70F")
	proposalContractAddress   = ethcommon.HexToAddress("0x47c0d7D6689b55875ae75B2CF41E75400B76927F")
	credentialContractAddress = ethcommon.HexToAddress("0x0eEf1dA46ca651F08a76dEEA49BcbAc218bD92c8")
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
