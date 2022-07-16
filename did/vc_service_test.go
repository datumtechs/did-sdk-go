package did

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/types/claim"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ethcontext chainclient.Context
var didService *DIDService

var did = "did:pid:123"
var doc_context = "http://datumtech.com/did/v1"
var pctId = 1
var issuer = "did:pid:ICAC"
var publicKey *ecdsa.PublicKey

func NewMockEthClientContext(chainUrl string, priKey *ecdsa.PrivateKey) *chainclient.EthContext {
	ctx := new(chainclient.EthContext)
	ctx.SetPrivateKey(priKey)
	ctx.SetWalletAddress(crypto.PubkeyToAddress(priKey.PublicKey))
	return ctx
}

func init() {
	fmt.Println("initing........")
	key, _ := crypto.GenerateKey()
	publicKey = &key.PublicKey
	ethcontext = NewMockEthClientContext("http://localhost:8485", key)
	didService = NewDIDService(ethcontext)
}

func Test_createVC(t *testing.T) {
	claimVar := make(claim.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"
	response := didService.CreateCredential(did, doc_context, pctId, claimVar, expirationDate, issuer)

	t.Logf("response.Status:%d", response.Status)
}

func Test_verifyVC(t *testing.T) {
	claimVar := make(claim.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"
	response := didService.CreateCredential(did, doc_context, pctId, claimVar, expirationDate, issuer)

	a := assert.New(t)
	a.Equal(response.Status, Response_SUCCESS)

	ok := didService.VerifyCredential((response.Data).Credential, nil, publicKey)
	a.Equal(ok, true)
}
