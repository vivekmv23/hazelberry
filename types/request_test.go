package types

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	testingutils "github.com/vivekmv23/hazelberry/utils"
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
	reader := strings.NewReader(validRequestString)
	decoder := json.NewDecoder(reader)
	var r []Request
	testingutils.FatalIfError(decoder.Decode(&r), t)
	for _, validReq := range r {
		err := validReq.InitAndValidate()
		testingutils.FatalIfError(err, t)
		testingutils.FatalIfNotEquals(http.MethodGet, validReq.Method, t)
		testingutils.FatalIfTrue("request is not empty", validReq.IsEmpty(), t)
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
	reader := strings.NewReader(invalidRequestsString)
	decoder := json.NewDecoder(reader)
	var r []Request
	err := decoder.Decode(&r)
	testingutils.FatalIfError(err, t)
	for _, invalidRequest := range r {
		err := invalidRequest.InitAndValidate()
		testingutils.FatalIfNoError(err, t)
	}
}

func TestRequest_isEmpty(t *testing.T) {
	r := Request{}
	testingutils.FatalIfFalse("request is empty", r.IsEmpty(), t)
}
