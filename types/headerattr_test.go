package types

import (
	"encoding/json"
	testutils "github.com/vivekmv23/hazelberry/utils"
	"strings"
	"testing"
)

const validHeaderString string = `
{

	"key": "Content-Type",
	"value": "application/json",
	"disabled": false		
	
}
`

func TestHeaderAttr_valid(t *testing.T) {
	dataReader := strings.NewReader(validHeaderString)
	decoder := json.NewDecoder(dataReader)
	var validHeader HeaderAttr
	err := decoder.Decode(&validHeader)
	testutils.FatalIfError(err, t)
	err = validHeader.InitAndValidate()
	testutils.FatalIfError(err, t)
	testutils.FatalIfNotEquals(validHeader.getKey(), "Content-Type", t)
	testutils.FatalIfNotEquals(validHeader.getValue(), "application/json", t)
}

const invalidHeaderAttrsString string = `
[
	{

	},
	{
		"value": "application/json"
	}
]
`

func TestHeaderAttr_invalids(t *testing.T) {
	var invalidHeaderAttrs []HeaderAttr
	dataReader := strings.NewReader(invalidHeaderAttrsString)
	decoder := json.NewDecoder(dataReader)
	testutils.FatalIfError(decoder.Decode(&invalidHeaderAttrs), t)
	for _, invalidHeaderAttr := range invalidHeaderAttrs {
		testutils.FatalIfNoError(invalidHeaderAttr.InitAndValidate(), t)
	}
}
