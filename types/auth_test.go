package types

import (
	"encoding/json"
	testutils "hazelberry/utils"
	"strings"
	"testing"
)

const valid string = `
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
	dataReader := strings.NewReader(valid)
	decoder := json.NewDecoder(dataReader)
	var validAuth Auth
	err := decoder.Decode(&validAuth)
	testutils.FatalIfMsg(testutils.NoError(err), t)
	err = validAuth.InitAndValidate()
	testutils.FatalIfMsg(testutils.NoError(err), t)
	testutils.FatalIfMsg(testutils.Equals(validAuth.getId(), "some-username"), t)
	testutils.FatalIfMsg(testutils.Equals(validAuth.getPass(), "some-password"), t)
}

const invalids string = `
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
	dataReader := strings.NewReader(invalids)
	decoder := json.NewDecoder(dataReader)
	testutils.FatalIfMsg(testutils.NoError(decoder.Decode(&invalidAuths)), t)
	for _, invalidAuth := range invalidAuths {
		err := invalidAuth.InitAndValidate()
		testutils.FatalIfMsg(testutils.HasError(err), t)
	}
}
