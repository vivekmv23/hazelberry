package types

import (
	"encoding/json"
	testutils "hazelberry/utils"
	"strings"
	"testing"
)

const validAuthString string = `
{
	"type": "basic",
	"basic": [
		{
			"key": "password",
			"value": "some-password",
			"type": "string"
		},
		{
			"key": "username",
			"value": "some-username",
			"type": "string"
		}
	]
}
`

func TestAuth_valid(t *testing.T) {
	dataReader := strings.NewReader(validAuthString)
	decoder := json.NewDecoder(dataReader)
	var validAuth Auth
	err := decoder.Decode(&validAuth)
	testutils.FatalIfError(err, t)
	err = validAuth.InitAndValidate()
	testutils.FatalIfError(err, t)
	testutils.FatalIfNotEquals(validAuth.getId(), "some-username", t)
	testutils.FatalIfNotEquals(validAuth.getPass(), "some-password", t)
}

const invalidAuthListString string = `
[
	{
		"type": "unsupported",
		"unsupported": []
	},
	{
		"basic": []
	},
	{
		"type": "basic",
		"unsupported": []
	},
	{
		"type": "basic",
		"basic": []
	},
	{
		"type": "basic",
		"basic": [
			{
				"key": "username",
				"value": "some-username",
				"type": "string"
			}
		]
	},
	{
		"type": "basic",
		"basic": [
			{
				"key": "password",
				"value": "some-password",
				"type": "string"
			}
		]
	},
	{
		"type": "basic",
		"basic": [
			{
				"key": "password"
			},
			{
				"key": "username"
			}
		]
	}
]`

func TestAuth_invalids(t *testing.T) {
	var invalidAuths []Auth
	dataReader := strings.NewReader(invalidAuthListString)
	decoder := json.NewDecoder(dataReader)
	testutils.FatalIfError(decoder.Decode(&invalidAuths), t)
	for _, invalidAuth := range invalidAuths {
		testutils.FatalIfNoError(invalidAuth.InitAndValidate(), t)
	}
}
