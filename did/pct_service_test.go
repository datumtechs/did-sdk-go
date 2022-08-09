package did

import (
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var schema = "{\"additionalProperties\":false,\"properties\":{\"nodeID\":{\"type\":\"string\"},\"nodeName\":{\"type\":\"string\"},\"url\":{\"type\":\"string\"}}}"

func Test_VerifyContentByJsonSchema(t *testing.T) {
	content := make(map[string]interface{})
	content["nodeID"] = "did:pid:testNodeID"
	content["nodeName"] = "testNodeName"
	content["url"] = "http://www.datumtechs.org"
	b := common.VerifyWithJsonSchema(schema, content)

	a := assert.New(t)
	a.Equal(true, b)
}

func Test_registerPct(t *testing.T) {
	t.Helper()
	setup()

	req := CreatePctReq{}
	req.PrivateKey = MockWalletInstance().GetPrivateKey()
	req.PctJson = "{\"additionalProperties\":false,\"properties\":{\"nodeID\":{\"type\":\"string\"},\"nodeName\":{\"type\":\"string\"},\"url\":{\"type\":\"string\"}}}"
	req.Extra = []byte{}
	response := didService.PctService.RegisterPct(req)
	t.Log(response)

	a := assert.New(t)

	if a.Equal(Response_SUCCESS, response.Status) {
		pctId, _ := new(big.Int).SetString(string(response.Data), 10)
		getResponse := didService.PctService.GetPct(pctId)
		a.Equal(req.PctJson, getResponse.Data.JsonSchema)
	}
}

func Test_GetPct(t *testing.T) {
	//t.Helper()
	setup()
	getResponse := didService.PctService.GetPct(new(big.Int).SetUint64(1000))

	t.Logf("pct schema: %+v", *getResponse.Data)
	a := assert.New(t)
	a.Equal(Response_SUCCESS, getResponse.Status)

}
