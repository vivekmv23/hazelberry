package types

import (
	"encoding/json"
	"strings"
	"testing"

	testingutil "github.com/vivekmv23/hazelberry/utils"
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

	}
]
`

func TestCollection_valids(t *testing.T) {
	reader := strings.NewReader(validCollectionString)
	decoder := json.NewDecoder(reader)
	var c []Collection
	testingutil.FatalIfError(decoder.Decode(&c), t)
	for _, validCollection := range c {
		err := validCollection.InitAndValidate()
		testingutil.FatalIfError(err, t)
		testingutil.FatalIfNotEquals("some-name", validCollection.Info.Name, t)
		for _, item := range validCollection.Item {
			testingutil.FatalIfNotEquals("https://some-url.com", item.Request.Url.Raw, t)
		}
		if !validCollection.Auth.IsEmpty() {
			testingutil.FatalIfNotEquals("some-username", validCollection.Auth.getId(), t)
			testingutil.FatalIfNotEquals("some-password", validCollection.Auth.getPass(), t)
		}
		testingutil.FatalIfTrue("collection is empty", validCollection.IsEmpty(), t)
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
	reader := strings.NewReader(invalidCollectionsString)
	decoder := json.NewDecoder(reader)
	var c []Collection
	testingutil.FatalIfError(decoder.Decode(&c), t)
	for _, invalidCollection := range c {
		testingutil.FatalIfNoError(invalidCollection.InitAndValidate(), t)
	}
}
