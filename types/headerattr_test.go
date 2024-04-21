package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
)

const validHeaderString string = `
{

	"key": "Content-Type",
	"value": "application/json",
	"disabled": false		
	
}
`

func TestHeaderAttr_valid(t *testing.T) {
	var validHeader HeaderAttr
	testutil.Decode(validHeaderString, &validHeader, t)
	err := validHeader.InitAndValidate()
	testutil.FatalIfError(err, t)
	testutil.FatalIfNotEquals(validHeader.getKey(), "Content-Type", t)
	testutil.FatalIfNotEquals(validHeader.getValue(), "application/json", t)
}

const invalidHeaderAttrsString string = `
[
	{},
	{
		"value": "application/json"
	}
]
`

func TestHeaderAttr_invalids(t *testing.T) {
	var ha []HeaderAttr
	testutil.Decode(invalidHeaderAttrsString, &ha, t)
	for _, invalidHeaderAttr := range ha {
		testutil.FatalIfNoError(invalidHeaderAttr.InitAndValidate(), t)
	}
}
