package did

import (
	"crypto/ecdsa"
	"fmt"
	platoncommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var didService *DIDService

var did = "did:pid:lat1x4w7852dxs69sy2mgf8w0s7tmvqx3cz2ydaxq4"
var doc_context = "http://datumtech.com/did/v1"
var pctId = new(big.Int).SetUint64(1)
var issuer = "did:pid:lat1x4w7852dxs69sy2mgf8w0s7tmvqx3cz2ydaxq4"
var publicKeyId string
var privateKey *ecdsa.PrivateKey

func setup() {
	fmt.Println("initing........")
	InitMockWallet()

	privateKey = MockWalletInstance().priKey
	publicKeyId = MockWalletInstance().walletAddress.Hex() + "#keys-1"

	ethcontext := chainclient.NewEthClientContext("", "lat", MockWalletInstance())
	didService = NewDIDService(ethcontext)
}

func Test_bech32(t *testing.T) {
	setup()
	addr := string(([]byte(did))[len("did:pid:"):])
	platonAddress, err := platoncommon.Bech32ToAddress(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(platonAddress)
}

func Test_createVC(t *testing.T) {
	setup()
	claimVar := make(types.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"
	req := new(CreateCredentialReq)
	req.Did = did
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId

	response := didService.VcService.CreateCredentialSimple(*req)

	t.Logf("response.Data:%+v", response.Data)

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)

}

func Test_verifyVC(t *testing.T) {
	setup()

	claimVar := make(types.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"

	req := new(CreateCredentialReq)
	req.Did = did
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId

	response := didService.VcService.CreateCredentialSimple(*req)
	t.Logf("response.Status:%d", response.Status)
	cred := response.Data
	t.Logf("%#v", cred)
	a := assert.New(t)
	if !a.Equal(Response_SUCCESS, response.Status) {
		t.Fatal()
	}

	ok := didService.VcService.VerifyVCWithPublicKey(&cred, didService.VcService.ctx.GetPublicKey())
	a.Equal(true, ok)
}
