package did

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreatePresentation(t *testing.T) {
	setup()

	claimVar := creteComplexClaim()

	disclosureVar := creteComplexDisclosure()

	req := new(CreatePresentationReq)
	req.Challenge = "1234567"
	req.PresentationPolicy = map[string]types.Claim{pctId.String(): disclosureVar}
	vc := createCredential(claimVar)
	req.Credential = []*types.Credential{vc}

	req.Authentication = Authentication{Issuer: vc.Issuer, IssuerPrivateKey: privateKey, PublicKeyId: publicKeyId}

	response := didService.CredentialService.CreatePresentation(*req)

	b, _ := json.Marshal(response.Data)

	t.Logf("response.Data:%s", string(b))

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)

}

func createCredential(claim types.Claim) *types.Credential {

	expirationDate := "1989-06-06-18T21:19:10"
	req := new(CreateCredentialReq)
	req.Did = did
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claim
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId
	req.Type = types.CREDENTIAL_TYPE_VC
	response := didService.CredentialService.CreateCredentialSimple(*req)
	return &response.Data
}

func creteSimpleClaim() types.Claim {
	claimVar := make(types.Claim)
	claimVar["name"] = "Alice"
	claimVar["age"] = "12"
	return claimVar
}

func creteSimpleDisclosure() types.Claim {
	claimVar := make(types.Claim)
	claimVar["name"] = 1
	claimVar["age"] = 0
	return claimVar
}

func creteComplexClaim() types.Claim {
	claimVar := make(types.Claim)
	claimVar["name"] = "Alice"
	claimVar["age"] = "12"

	degreeMap := make(map[string]interface{})
	degreeMap["art"] = "A+"
	degreeMap["math"] = "A-"
	degreeMap["chemistry"] = "C"
	degreeMap["physics"] = "A"
	degreeMap["language"] = "A+"

	claimVar["degree"] = degreeMap
	return claimVar
}

func creteComplexDisclosure() types.Claim {
	claimVar := make(types.Claim)
	claimVar["name"] = 1
	claimVar["age"] = 0

	degreeMap := make(map[string]interface{})
	degreeMap["art"] = 1
	degreeMap["math"] = 1
	degreeMap["chemistry"] = 0
	degreeMap["physics"] = 1
	degreeMap["language"] = 1

	claimVar["degree"] = degreeMap
	return claimVar
}
