package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
)

var validCollectionString string = `
[
	{
		"info": {
			"name": "some-name"
		},
		"item": [
			{
				"name": "sample-item-1",
				"id": "sample-id-1",
				"request": {
					"method": "GET",
					"url": "https://some-url.com"
				}
			},
			{
				
				"name": "sample-item-2",
				"id": "sample-id-2",
				"request": {
					"method": "DELETE",
					"url": "https://some-url.com"
				}

			}
		],
		"auth": {
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

	},
	{
		"info": {
			"name": "some-name"
		},
		"item": [
			{
				"name": "sample-item",
				"id": "sample-id",
				"request": {
					"method": "GET",
					"url": "https://some-url.com"
				}
			}
		]

	},
	{
		"info": {
			"name": "some-name"
		},
		"item": [
			{
				"name": "sample-item",
				"id": "sample-id",
				"request": {
					"method": "GET",
					"url": "https://some-url.com"
				}
			}
		],
		"variable": [
			{
				"key": "some-var-key",
				"value": "some-var-value"
			}
    	]
	}
]
`

func TestCollection_valids(t *testing.T) {
	var c []Collection
	testutil.Decode(validCollectionString, &c, t)
	for _, validCollection := range c {
		err := validCollection.InitAndValidate()
		testutil.FatalIfError(err, t)
		testutil.FatalIfNotEquals("some-name", validCollection.Info.Name, t)
		for _, item := range validCollection.Item {
			testutil.FatalIfNotEquals("https://some-url.com", item.Request.Url.Raw, t)
		}
		for _, varaible := range validCollection.Varaible {
			testutil.FatalIfNotEquals("some-var-key", varaible.GetKey(), t)
			testutil.FatalIfNotEquals("some-var-value", varaible.GetValue(), t)
		}
		if !validCollection.Auth.IsEmpty() {
			testutil.FatalIfNotEquals("some-username", validCollection.Auth.GetId(), t)
			testutil.FatalIfNotEquals("some-password", validCollection.Auth.GetPass(), t)
		}
		testutil.FatalIfTrue("collection is empty", validCollection.IsEmpty(), t)
	}
}

var invalidCollectionsString string = `
[
	{},
	{
		"info": {
			"name": "some-name"
		}
	},
	{
		"info": {
			"name": "some-name"
		},
		"item": [
			{
				"name": "sample-item",
				"id": "sample-id",
				"request": {
					"method": "GET",
					"url": "https://some-url.com"
				}
			},
			{
				"name": "sample-item",
				"id": "sample-id"
			}
		]
	},
	{
		"info": {
			"name": "some-name"
		},
		"item": [
			{
				"name": "sample-item",
				"id": "sample-id",
				"request": {
					"method": "GET",
					"url": "https://some-url.com"
				}
			}
		],
		"variable":[{}]
	},
	{
		"info": {
			"name": "some-name"
		},
		"item": [
			{
				"name": "sample-item",
				"id": "sample-id",
				"request": {
					"method": "GET",
					"url": "https://some-url.com"
				}
			}
		],
		"auth": {
			"type": "UNSUPPORTED",
			"basic": [
				{
					"key": "password",
					"value": "some-password",
					"type": "string"
				}
			]
		}
	}
]
`

func TestCollection_invalids(t *testing.T) {
	var c []Collection
	testutil.Decode(invalidCollectionsString, &c, t)
	for _, invalidCollection := range c {
		testutil.FatalIfNoError(invalidCollection.InitAndValidate(), t)
	}
}
