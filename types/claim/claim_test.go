package claim

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_GenerateClaimRandSalt(t *testing.T) {
	jsonString := "{\"nodeID\":\"testNodeID\",\"nodeName\":\"testNodeName\",\"url\":\"http://www.datumtechs.org\"}"

	claim := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &claim)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	hashStringBuilder := strings.Builder{}
	seed := common.Uint64ToBigEndianBytes(uint64(23523865082340324))
	GenerateClaimSaltForMap(claim, seed, &hashStringBuilder)
	t.Log(claim)
	t.Log(common.BigEndianBytesToUint64(seed))
	t.Log(hashStringBuilder.String())
}

func Test_marshalJson(t *testing.T) {
	v := 3.1415926535
	b, _ := json.Marshal(v)
	a := assert.New(t)
	a.Equal("3.1415926535", string(b))
}
