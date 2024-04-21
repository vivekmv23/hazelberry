package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
)

var validBodyString string = `
{
	"mode": "raw",
	"raw": "{\r\n    \"some-body-key\": \"some-body-value\"\r\n}"
}
`

func TestBody_valid(t *testing.T) {
	var b Body
	testutil.Decode(validBodyString, &b, t)
	err := b.InitAndValidate()
	testutil.FatalIfError(err, t)
	testutil.FatalIfNotEquals("raw", b.Mode, t)
	testutil.FatalIfNotEquals("{\r\n    \"some-body-key\": \"some-body-value\"\r\n}", b.Raw, t)
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
	var b []Body
	testutil.Decode(invalidBodyString, &b, t)
	for _, invalidBody := range b {
		err := invalidBody.InitAndValidate()
		testutil.FatalIfNoError(err, t)
	}

}

func TestBody_isEmpty(t *testing.T) {
	emptyBody := Body{}
	testutil.FatalIfFalse("Empty body", emptyBody.IsEmpty(), t)
}
