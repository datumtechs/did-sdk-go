package santhosh_tekuri

import (
	"encoding/json"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_moreKeys(t *testing.T) {

	sch, err := jsonschema.Compile("../pctschema/datum-schema.json")
	if err != nil {
		t.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile("../pctschema/datum-instance-moreKeys.json")
	if err != nil {
		t.Fatal(err)
	}

	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatal(err)
	}

	a := assert.New(t)

	if err = sch.Validate(v); err != nil {
		//t.Fatalf("%#v", err)
		a.ErrorContains(err, "additionalProperties 'address' not allowed")
	}

}
