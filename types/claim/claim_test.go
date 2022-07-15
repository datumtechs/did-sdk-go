package claim

import (
	"encoding/json"
	"testing"
)

func Test_GenerateClaimFixedSalt(t *testing.T) {
	jsonString := "{\"id\":\"did:example:ebfeb1f712ebc6f1c276e12ec21\",\"alumniOf\":{\"id\":\"did:example:c276e12ec21ebfeb1f712ebc6f1\",\"name\":[{\"value\":\"Example University\",\"lang\":\"en\"},{\"value\":\"Exemple d'Université\",\"lang\":\"fr\"}]}}"

	claim := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &claim)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	GenerateClaimSaltForMap(claim, "abc123")
	t.Log(claim)

}

func Test_GenerateClaimRandSalt(t *testing.T) {
	jsonString := "{\"id\":\"did:example:ebfeb1f712ebc6f1c276e12ec21\",\"alumniOf\":{\"id\":\"did:example:c276e12ec21ebfeb1f712ebc6f1\",\"name\":[{\"value\":\"Example University\",\"lang\":\"en\"},{\"value\":\"Exemple d'Université\",\"lang\":\"fr\"}]}}"

	claim := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &claim)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	GenerateClaimSaltForMap(claim, "")
	t.Log(claim)

}
