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
    "host": "some-url.some-domain.com"
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
  },
  {
    "raw": "https://some-url.some-domain.com?some-param-key=some-param-value",
    "protocol": "https",
    "variable": [
      {
        "key": "some-var-key",
        "value": "some-var-value"
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
		for _, q := range validUrls.Query {
			testutil.FatalIfNotEquals("some-param-key", q.GetKey(), t)
			testutil.FatalIfNotEquals("some-param-value", q.GetValue(), t)
			testutil.FatalIfFalse("disabled", q.IsDisabled(), t)
		}
		for _, v := range validUrls.Variable {
			testutil.FatalIfNotEquals("some-var-key", v.GetKey(), t)
			testutil.FatalIfNotEquals("some-var-value", v.GetValue(), t)
			testutil.FatalIfTrue("disabled", v.IsDisabled(), t)
		}
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
        "key" : "some-param-key",
        "value": "some-param-value"
      },
      {}
    ]
  },
  {
    "raw": "https://some-url.some-domain.com?some-param-key=some-param-value",
    "protocol": "https",
    "host": "some-url.some-domain.com",
    "variable": [
      {}
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
