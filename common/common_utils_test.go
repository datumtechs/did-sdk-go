package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
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
