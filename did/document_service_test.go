package did

import (
	"github.com/datumtechs/did-sdk-go/types"
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
	response := didService.DocumentService.QueryDidDocument(applicantDid)
	t.Logf("response:%+v", *response.Data)
}
