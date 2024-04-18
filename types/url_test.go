package types

import (
	"encoding/json"
	testutils "github.com/vivekmv23/hazelberry/utils"
	"strings"
	"testing"
)

var validUrlsString string = `
[
  {
    "raw": "https://some-url.some-domain.com?some-param-key=some-param-value",
    "protocol": "https",
    "host": [
      "some-url",
      "some-domain",
      "com"
    ],
    "query": [
      {
        "key": "some-param-key",
        "value": "some-param-value",
        "disabled": true
      }
    ]
  },
  {
    "raw": "https://some-url.some-domain.com?some-param-key=some-param-value",
    "protocol": "https",
    "host": "some-url.some-domain.com",
    "query": [
      {
        "key": "some-param-key",
        "value": "some-param-value",
        "disabled": true
      }
    ]
  },
  {
    "raw": "https://some-url.some-domain.com?some-param-key=some-param-value",
    "protocol": "https",
    "query": [
      {
        "key": "some-param-key",
        "value": "some-param-value",
        "disabled": true
      }
    ]
  }
]
`

func TestUrl_valid(t *testing.T) {
	reader := strings.NewReader(validUrlsString)
	decoder := json.NewDecoder(reader)
	var u []Url
	testutils.FatalIfError(decoder.Decode(&u), t)
	for _, validUrls := range u {
		err := validUrls.InitAndValidate()
		testutils.FatalIfError(err, t)
		testutils.FatalIfNotEquals("https://some-url.some-domain.com?some-param-key=some-param-value", validUrls.Raw, t)
		hostString, err := validUrls.GetHost()
		testutils.FatalIfError(err, t)
		if validUrls.Host != nil {
			testutils.FatalIfNotEquals("some-url.some-domain.com", hostString, t)
		}
		testutils.FatalIfFalse("query paraments length is one", len(validUrls.Query) == 1, t)
		testutils.FatalIfNotEquals("some-param-key", validUrls.Query[0].getKey(), t)
		testutils.FatalIfNotEquals("some-param-value", validUrls.Query[0].getValue(), t)
		testutils.FatalIfFalse("disabled", validUrls.Query[0].Disabled, t)
	}
}

var invalidUrlsString string = `
[
  {},
  {
    "raw": "https://some-url.some-domain.com?some-param-key=some-param-value",
    "protocol": "https",
    "host": false
  },
  {
    "raw": "https://some-url.some-domain.com?some-param-key=some-param-value",
    "protocol": "https",
    "host": "some-url.some-domain.com",
    "query": [
      {
        "value": "some-param-value"
      }
    ]
  }
]
`

func TestUrl_invalid(t *testing.T) {
	reader := strings.NewReader(invalidUrlsString)
	decoder := json.NewDecoder(reader)
	var u []Url
	testutils.FatalIfError(decoder.Decode(&u), t)
	for _, invalidUrls := range u {
		err := invalidUrls.InitAndValidate()
		testutils.FatalIfNoError(err, t)
	}
}

func TestUrl_isEmpty(t *testing.T) {
	emptyUrl := Url{}
	testutils.FatalIfFalse("Empty url", emptyUrl.IsEmpty(), t)
}
