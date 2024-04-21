package types

import (
	"encoding/json"
	"strings"
	"testing"

	testingutils "github.com/vivekmv23/hazelberry/utils"
)

var validInfoString string = `
{
	"name": "sample-name"
}
`

func TestInfo_valid(t *testing.T) {
	reader := strings.NewReader(validInfoString)
	decoder := json.NewDecoder(reader)
	var inf Info
	testingutils.FatalIfError(decoder.Decode(&inf), t)
	err := inf.InitAndValidate()
	testingutils.FatalIfError(err, t)
	testingutils.FatalIfNotEquals("sample-name", inf.Name, t)
	testingutils.FatalIfTrue("info name empty", inf.IsEmpty(), t)
}

var invalidInfoString string = "{}"

func TestInfo_invalid(t *testing.T) {
	reader := strings.NewReader(invalidInfoString)
	decoder := json.NewDecoder(reader)
	var inf Info
	testingutils.FatalIfError(decoder.Decode(&inf), t)
	err := inf.InitAndValidate()
	testingutils.FatalIfNoError(err, t)
	testingutils.FatalIfFalse("info name empty", inf.IsEmpty(), t)
}
