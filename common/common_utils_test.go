package common

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_json_string(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := RandStringUnsafe(8)
		t.Logf("rand string: %s", r)
	}
}

func Test_cloneClaim(t *testing.T) {
	jsonString := "{\"id\":\"did:example:ebfeb1f712ebc6f1c276e12ec21\",\"alumniOf\":{\"id\":\"did:example:c276e12ec21ebfeb1f712ebc6f1\",\"name\":[{\"value\":\"Example University\",\"lang\":\"en\"},{\"value\":\"Exemple d'Université\",\"lang\":\"fr\"}]}}"
	src := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &src)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	dest := Clone(src)
	dest["id"] = "newId"

	t.Logf("src %v", src)
	t.Logf("dest: %v", dest)

	a := assert.New(t)
	a.Equal("did:example:ebfeb1f712ebc6f1c276e12ec21", src["id"])
	a.Equal("newId", dest["id"])

}

var seed = uint64(23523865082340324)

// salt[0]=3630861241102706729 salt[1]=12004035919866408515 salt[2]=16714233546749475040 salt[3]=12244885108623859682 salt[4]=17247414944566619985 salt[5]=1579923115120900640 salt[6]=852460680154286706 salt[7]=14713654450971305745 salt[8]=3123376975702223225 salt[9]=8786413226062047404
func Test_GenerateSequence256(t *testing.T) {
	for i := 0; i < 5; i++ {
		result := GenerateSequence256(Uint64ToBigEndianBytes(seed), 10)
		for idx, seed := range result {
			fmt.Printf("salt[%d]=%d ", idx, BigEndianBytesToUint64(seed))
		}
		fmt.Print("\n")
	}
}

func Test_VerifyValidJsonSchema(t *testing.T) {
	validSchema := "{\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"nodeID\": {\n      \"type\": \"string\"\n    },\n    \"nodeName\": {\n      \"type\": \"string\"\n    },\n    \"url\": {\n      \"type\": \"string\"\n    }\n  }\n}"

	result := VerifyJsonSchema(validSchema)

	a := assert.New(t)
	a.Equal(true, result)
}
func Test_VerifyInvalidJsonSchema(t *testing.T) {
	invalidSchema := "{\n  \"additionalProperties\": false,\n  \"properties\": {\n    \"nodeID\": {\n      \"type\": \"string\"\n    },\n    \"nodeName\": {\n      \"type\": \"string\"\n    },\n    \"url\": {\n      \"type\": \"string\",\n    }\n  }\n}"

	result := VerifyJsonSchema(invalidSchema)

	a := assert.New(t)
	a.Equal(false, result)
}

func Test_utc(t *testing.T) {
	t.Log(FormatUTC(time.Now()))
	t.Log(FormatUTC(time.Now().UTC()))
	t.Log(time.Now().Format(time.RFC3339))

	TIME, _ := time.ParseInLocation(time.RFC3339, "2022-08-24 15:41:08 +0800", time.UTC)
	t.Log(TIME)

}
