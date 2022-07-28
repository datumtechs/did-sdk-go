package did

import (
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/stretchr/testify/assert"
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
