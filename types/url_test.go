package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
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
	var u []Url
	testutil.Decode(validUrlsString, &u, t)
	for _, validUrls := range u {
		err := validUrls.InitAndValidate()
		testutil.FatalIfError(err, t)
		testutil.FatalIfNotEquals("https://some-url.some-domain.com?some-param-key=some-param-value", validUrls.Raw, t)
		hostString, err := validUrls.GetHost()
		testutil.FatalIfError(err, t)
		if validUrls.Host != nil {
			testutil.FatalIfNotEquals("some-url.some-domain.com", hostString, t)
		}
		testutil.FatalIfFalse("query paraments length is one", len(validUrls.Query) == 1, t)
		testutil.FatalIfNotEquals("some-param-key", validUrls.Query[0].getKey(), t)
		testutil.FatalIfNotEquals("some-param-value", validUrls.Query[0].getValue(), t)
		testutil.FatalIfFalse("disabled", validUrls.Query[0].Disabled, t)
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
	var u []Url
	testutil.Decode(invalidUrlsString, &u, t)
	for _, invalidUrls := range u {
		err := invalidUrls.InitAndValidate()
		testutil.FatalIfNoError(err, t)
	}
}

func TestUrl_isEmpty(t *testing.T) {
	emptyUrl := Url{}
	testutil.FatalIfFalse("Empty url", emptyUrl.IsEmpty(), t)
}
