package did

import (
	platoncommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createIssuerDid(t *testing.T) {
	setup()
	req := CreateDidReq{}
	req.PrivateKey = privateKey
	req.PublicKey = publicKey
	req.PublicKeyType = types.PublicKey_SECP256K1
	response := didService.DocumentService.CreateDID(req)
	t.Logf("response.data:%+v", response.Data)

	a := assert.New(t)
	if a.Contains(response.Msg, "Did exists") || a.Equal(Response_SUCCESS, response.Status) {
		docResponse := didService.DocumentService.QueryDidDocument(response.Data)
		t.Logf("response.Data:%+v", *docResponse.Data)
		t.Logf("pubkey:%+v", *docResponse.Data.PublicKey[0])
	}
}

func Test_createApplicantDid(t *testing.T) {
	setup()
	req := CreateDidReq{}
	req.PrivateKey = applicantPriKey
	req.PublicKey = applicantPublicKey
	req.PublicKeyType = types.PublicKey_SECP256K1
	response := didService.DocumentService.CreateDID(req)
	t.Logf("response.data:%+v", response.Data)

	a := assert.New(t)
	if a.Contains(response.Msg, "Did exists") || a.Equal(Response_SUCCESS, response.Status) {
		docResponse := didService.DocumentService.QueryDidDocument(response.Data)
		t.Logf("response.Data:%+v", *docResponse.Data)
		t.Logf("pubkey:%+v", *docResponse.Data.PublicKey[0])
	}
}

func Test_QueryDidDocument(t *testing.T) {
	setup()
	response := didService.DocumentService.QueryDidDocument("did:pid:lat1g83xnwcqc4uufpx9588muq0e98g93rlmpgx563")
	t.Logf("response:%+v", *response.Data.PublicKey[0])
}

func Test_HexToPublicKey(t *testing.T) {
	setup()

	pubKey := crypto.HexToPublicKey(ethhexutil.Encode(ethcommon.FromHex("0x049e5d1a3e18606fb9aaee2d4e28a2098a0a78525d12a1e9fdfaaf930adba770a3d9026f34df5e94053f88dad9efef98c7bc731d46212fb46b3b98d86e09788da0")))
	t.Logf("pubKey:%+v", pubKey)

	bech32Addr := platoncommon.Address(ethcrypto.PubkeyToAddress(*pubKey))

	t.Logf("bech32Addr:%+v", bech32Addr.Bech32())

}
