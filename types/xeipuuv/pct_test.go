package xeipuuv

import (
	"encoding/json"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"testing"
)

func Test_inValidSchema(t *testing.T) {

	schemaLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-schema-invalid.json")
	documentLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-instance.json")

	schemaLoader.JsonReference()
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func Test_validByPct(t *testing.T) {

	schemaLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-instance.json")
	//stringLoader := gojsonschema.NewStringLoader("{\"nodeID\":{\"type\":\"did:example:ebfeb1f712ebc6f1c276e12ec21\"},\"nodeName\":{\"type\":\"init Node\"},\"url\":{\"type\":\"https://www.initnode.org\"}}")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func Test_complexValue(t *testing.T) {

	schemaLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-schema.json")
	//documentLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-instance.json")
	stringLoader := gojsonschema.NewStringLoader("{\"nodeID\":\"did:example:c276e12ec21ebfeb1f712ebc6f1\",\"nodeName\":\"[{\\\"value\\\":\\\"Example University\\\",\\\"lang\\\":\\\"en\\\"},{\\\"value\\\":\\\"Exemple d'Université\\\",\\\"lang\\\":\\\"fr\\\"}]\",\"url\":\"https://www.darumtechs.org\"}")

	result, err := gojsonschema.Validate(schemaLoader, stringLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func Test_moreKeys(t *testing.T) {

	schemaLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-instance-moreKeys.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func Test_lessKeys(t *testing.T) {

	schemaLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file://D:/github.com/datumtechs/did-sdk-go/types/pctschema/datum-instance-lessKeys.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func Test_jsonSchema(t *testing.T) {
	schema := "{\"properties\":{\"nodeID\":{\"type\":\"string\"},\"nodeName\":{\"type\":\"string\"},\"url\":{\"type\":\"string\"}}}"
	var m map[string]interface{}
	json.Unmarshal([]byte(schema), &m)
	fmt.Print(m)
}

func Test_jsonSchema2(t *testing.T) {
	schema := "{\"nodeID\":\"did:example:c276e12ec21ebfeb1f712ebc6f1\",\"nodeName\":\"[{\\\"value\\\":\\\"Example University\\\",\\\"lang\\\":\\\"en\\\"},{\\\"value\\\":\\\"Exemple d'Université\\\",\\\"lang\\\":\\\"fr\\\"}]\",\"url\":\"https://www.darumtechs.org\"}"
	var m map[string]interface{}
	json.Unmarshal([]byte(schema), &m)
	fmt.Print(m)
}
