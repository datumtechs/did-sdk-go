package did

import (
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	if a.Contains(response.Msg, "Did exists already") || a.Equal(Response_SUCCESS, response.Status) {
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
	if a.Contains(response.Msg, "Did exists already") || a.Equal(Response_SUCCESS, response.Status) {
		docResponse := didService.DocumentService.QueryDidDocument(response.Data)
		t.Logf("response.Data:%+v", *docResponse.Data)
		t.Logf("pubkey:%+v", *docResponse.Data.PublicKey[0])
	}
}

func Test_QueryDidDocument(t *testing.T) {
	setup()
	response := didService.DocumentService.QueryDidDocument("did:pid:lat1uejmcfntvvpx7ddw4a28wpzr4puh8kl9248fln")
	t.Logf("response:%+v", *response.Data.PublicKey[0])
}

func Test_HexToPublicKey(t *testing.T) {
	setup()

	response := crypto.HexToPublicKey(hexutil.Encode(ethcommon.FromHex("0x042964fb72636ae2695dee6543dcaa2eaa67d3f8db86d3e5af1b1b059be45923466d3bb85e75c020d8079b6a7210bf33b22e7ec31d449914e52595253f41686ec9")))
	t.Logf("response:%+v", response)
}
