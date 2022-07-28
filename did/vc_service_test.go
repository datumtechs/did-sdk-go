package did

import (
	"fmt"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/types/claim"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var didService *DIDService

var did = "did:pid:123"
var doc_context = "http://datumtech.com/did/v1"
var pctId = new(big.Int).SetUint64(1)
var issuer = "did:pid:ICAC"

func setup() {
	fmt.Println("initing........")
	InitMockWallet()
	ethcontext := chainclient.NewEthClientContext("", MockWalletInstance())
	didService = NewDIDService(ethcontext)
}

func Test_createVC(t *testing.T) {
	setup()
	claimVar := make(claim.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"
	response := didService.vcService.CreateCredentialNotValidateClaim(did, doc_context, pctId, claimVar, expirationDate, issuer)

	t.Logf("response.Status:%d", response.Status)
}

func Test_verifyVC(t *testing.T) {
	setup()

	claimVar := make(claim.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"
	response := didService.vcService.CreateCredentialNotValidateClaim(did, doc_context, pctId, claimVar, expirationDate, issuer)
	cwrapper := response.Data
	t.Logf("%#v", *cwrapper.Credential)
	a := assert.New(t)
	a.Equal(response.Status, Response_SUCCESS)

	ok := didService.vcService.VerifyCredential((response.Data).Credential, nil, didService.vcService.ctx.GetPublicKey())
	a.Equal(ok, true)
}
