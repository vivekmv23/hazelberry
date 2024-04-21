package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
)

var validInfoString string = `
{
	"name": "sample-name"
}
`

func TestInfo_valid(t *testing.T) {
	var inf Info
	testutil.Decode(validInfoString, &inf, t)
	err := inf.InitAndValidate()
	testutil.FatalIfError(err, t)
	testutil.FatalIfNotEquals("sample-name", inf.Name, t)
	testutil.FatalIfTrue("info name empty", inf.IsEmpty(), t)
}

func TestInfo_invalid(t *testing.T) {
	var inf Info
	testutil.Decode("{}", &inf, t)
	err := inf.InitAndValidate()
	testutil.FatalIfNoError(err, t)
	testutil.FatalIfFalse("info name empty", inf.IsEmpty(), t)
}
