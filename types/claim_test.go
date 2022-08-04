package types

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

func Test_SplitClaim(t *testing.T) {
	jsonString := "{\"nodeID\":\"testNodeID\",\"nodeName\":\"testNodeName\",\"url\":\"http://www.datumtechs.org\"}"

	claim := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &claim)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	disclosureMap := common.Clone(claim)
	disclosureMap["nodeID"] = 0
	disclosureMap["nodeName"] = 1
	disclosureMap["url"] = 0

	undisclosedClaim := common.Clone(claim)

	seed := common.Uint64ToBigEndianBytes(uint64(23523865082340324))

	hashStringBuilder := strings.Builder{}
	GenerateClaimSaltForMap(claim, seed, &hashStringBuilder)
	t.Log(hashStringBuilder.String())

	seed = common.Uint64ToBigEndianBytes(uint64(23523865082340324))
	SplitForMap(claim, disclosureMap, undisclosedClaim, seed)
	t.Log(claim)
	t.Log(disclosureMap)
	t.Log(undisclosedClaim)

}

func Test_marshalJson(t *testing.T) {
	v := 3.1415926535
	b, _ := json.Marshal(v)
	a := assert.New(t)
	a.Equal("3.1415926535", string(b))
}

func Test_unmarshalClaim(t *testing.T) {
	//jsonString := "{\"nodeID\":\"testNodeID\",\"nodeName\":\"testNodeName\",\"url\":\"http://www.datumtechs.org\"}"

	jsonString := "{\"id\":\"did:cid:511112200001010015\",\"name\":\"小明\",\"alumniOf\":{\"id\":\"did:cedu:uestc\",\"name\":[{\"value\":\"电子科技大学\",\"lang\":\"cn\"}]},\"degree\":\"硕士研究生\",\"degreeType\":\"工科\",\"college\":\"计算机学院\"}"
	var claimMap Claim
	err := json.Unmarshal([]byte(jsonString), &claimMap)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	a := assert.New(t)
	t.Log(claimMap)
	a.Equal("小明", claimMap["name"])
}
