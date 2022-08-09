package did

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateEvidence(t *testing.T) {
	setup()
	t.Helper()
	req := CreateEvidenceReq{}
	req.PrivateKey = privateKey
	req.Credential = credential

	response := didService.CredentialService.CreateEvidence(req)

	t.Logf("response:%+v", response)

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	a.Equal(credentialHash.Hex(), response.Data)
}

func Test_QueryEvidence(t *testing.T) {
	setup()
	t.Helper()
	req := QueryEvidenceReq{}
	req.EvidenceId = "0x5d1c7cf3c760c1e8470c93055cd35fc2502059d668419efd38be6b37272583ce"

	response := didService.CredentialService.QueryEvidence(req)

	t.Logf("response:%+v", response)

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)
	a.Equal(publicKey, response.Data.SignerPubKey)
}
