package types

import (
	"net/http"
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
)

var validRequestString string = `
[
	{
		"auth": {
			"type": "basic",
			"basic": [
				{
					"key": "password",
					"value": "some-password"
				},
				{
					"key": "username",
					"value": "some-username"
				}
			]
		},
		"method": "GET",
		"header": [
			{
				"key": "some-header",
				"value": "some-header-value"
			}
		],
		"body": {
			"mode": "raw",
			"raw": "{\r\n    \"some-body-key\": \"some-body-value\"\r\n}"
		},
		"url": {
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
					"value": "some-param-value"
				}
			]
		}
	},
	{
		"method": "GET",
		"url": {
			"raw": "https://some-url.some-domain.com?some-param-key=some-param-value"
		}
	},
	{
		"method": "GET",
		"url": "https://some-url.some-domain.com?some-param-key=some-param-value"
	}
]
`

func TestRequest_valid(t *testing.T) {
	var r []Request
	testutil.Decode(validRequestString, &r, t)
	for _, validReq := range r {
		err := validReq.InitAndValidate()
		testutil.FatalIfError(err, t)
		testutil.FatalIfNotEquals(http.MethodGet, validReq.Method, t)
		testutil.FatalIfTrue("request is not empty", validReq.IsEmpty(), t)
		// Other items have been asserted in their own individual tests
	}

}

var invalidRequestsString string = `
[
	{},
	{
		"url": "https://some-url",
		"method": "UNSUPPORTED"
	},
	{
		"url": "https://some-url",
		"method": "GET",
		"auth": {
			"type": "basic",
			"basic": [
				{
					"key": "username",
					"value": "some-username"
				}
			]
		}
	},
	{
		"url": "https://some-url",
		"method": "GET",
		"body": {
			"mode": "UNSUPPORTED"
		}
	},
	{
		"url": "https://some-url",
		"method": "GET",
		"header": [
			{
				"value": "some-header-value"
			}
		]
	},
	{
		"url": 123445,
		"method": "UNSUPPORTED"
	}
]
`

func TestRequest_invalids(t *testing.T) {
	var r []Request
	testutil.Decode(invalidRequestsString, &r, t)
	for _, invalidRequest := range r {
		err := invalidRequest.InitAndValidate()
		testutil.FatalIfNoError(err, t)
	}
}

func TestRequest_isEmpty(t *testing.T) {
	r := Request{}
	testutil.FatalIfFalse("request is empty", r.IsEmpty(), t)
}

func TestRequest_urlParserError(t *testing.T) {
	req := Request{
		UrlParsed: make(chan int),
	}
	testutil.FatalIfNoError(convertParsedUrl(&req), t)
}
