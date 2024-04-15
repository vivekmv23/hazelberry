package types

import (
	"encoding/json"
	testingutils "github.com/vivekmv23/hazelberry/utils"
	"strings"
	"testing"
)

var validBodyString string = `
{
	"mode": "raw",
	"raw": "{\r\n    \"some-body-key\": \"some-body-value\"\r\n}"
}
`

func TestBody_valid(t *testing.T) {
	reader := strings.NewReader(validBodyString)
	decoder := json.NewDecoder(reader)
	var b Body
	testingutils.FatalIfError(decoder.Decode(&b), t)
	err := b.InitAndValidate()
	testingutils.FatalIfError(err, t)
	testingutils.FatalIfNotEquals("raw", b.Mode, t)
	testingutils.FatalIfNotEquals("{\r\n    \"some-body-key\": \"some-body-value\"\r\n}", b.Raw, t)
}

var invalidBodyString string = `
[
  {},
  {
    "mode": "graphql"
  }
]
`

func TestBody_invalid(t *testing.T) {
	reader := strings.NewReader(invalidBodyString)
	decoder := json.NewDecoder(reader)
	var b []Body
	testingutils.FatalIfError(decoder.Decode(&b), t)
	for _, invalidBody := range b {
		err := invalidBody.InitAndValidate()
		testingutils.FatalIfNoError(err, t)
	}

}
