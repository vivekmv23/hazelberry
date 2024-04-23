package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
)

var validVariableString string = `
[
	{
		"id": "some-id",
		"value": "some-value"
	},
	{
		"key": "some-id",
		"value": "some-value"
	}
]
`

func TestVariable_valids(t *testing.T) {
	var v []Variable
	testutil.Decode(validVariableString, &v, t)
	for _, validVar := range v {
		testutil.FatalIfError(validVar.InitAndValidate(), t)
		testutil.FatalIfTrue("variable not disabled", validVar.IsDisabled(), t)
		testutil.FatalIfNotEquals("some-id", validVar.GetKey(), t)
		testutil.FatalIfNotEquals("some-value", validVar.GetValue(), t)
	}
}

func TestVariable_invalid(t *testing.T) {
	invalidVar := Variable{}
	testutil.FatalIfNoError(invalidVar.InitAndValidate(), t)
}
