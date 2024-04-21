package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
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
	var validAuth Auth
	testutil.Decode(validAuthString, &validAuth, t)
	err := validAuth.InitAndValidate()
	testutil.FatalIfError(err, t)
	testutil.FatalIfNotEquals(validAuth.getId(), "some-username", t)
	testutil.FatalIfNotEquals(validAuth.getPass(), "some-password", t)
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
	testutil.Decode(invalidAuthListString, &invalidAuths, t)
	for _, invalidAuth := range invalidAuths {
		testutil.FatalIfNoError(invalidAuth.InitAndValidate(), t)
	}
}

var noauthAuthString = `
{
	"type": "noauth"
}
`

func TestAuth_isEmpty(t *testing.T) {
	a := Auth{}
	testutil.FatalIfFalse("Empty auth", a.IsEmpty(), t)
	testutil.Decode(noauthAuthString, &a, t)
	testutil.FatalIfError(a.InitAndValidate(), t)
	testutil.FatalIfFalse("Empty auth", a.IsEmpty(), t)
}
