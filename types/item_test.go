package types

import (
	"testing"

	"github.com/vivekmv23/hazelberry/testutil"
)

var validItemsString string = `
[
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
		"id": "sample-id",
		"request": {
			"method": "GET",
			"url": "https://some-url.com"
		},
		"variable": [
			{
				"id": "some-var-id",
				"value": "some-var-value"
			}
		]
	}
]
`

func TestItem_valid(t *testing.T) {
	var itms []Item
	testutil.Decode(validItemsString, &itms, t)
	for _, itm := range itms {
		err := itm.InitAndValidate()
		testutil.FatalIfError(err, t)
		testutil.FatalIfTrue("Item is empty", itm.IsEmpty(), t)
		testutil.FatalIfNotEquals("sample-item", itm.Name, t)
		testutil.FatalIfNotEquals("sample-id", itm.Id, t)
		testutil.FatalIfNotEquals("https://some-url.com", itm.Request.Url.Raw, t)
		for _, vari := range itm.Variable {
			testutil.FatalIfNotEquals("some-var-id", vari.GetKey(), t)
			testutil.FatalIfNotEquals("some-var-value", vari.GetValue(), t)
		}
		testutil.FatalIfFalse("item request auth is empty", itm.Request.Auth.IsEmpty(), t)
		testutil.FatalIfFalse("item request body is empty", itm.Request.Body.IsEmpty(), t)
	}

}

var invalidItemsString string = `
[
	{
		"request": {
			"method": "GET",
			"url": "https://some-url.com"
		}
	},
	{
		"name": "sample-item"
	},
	{
		"name": "sample-item",
		"request": {
			"method": "UNSUPPORTED",
			"url": "https://some-url.com"
		}
	},
	{
		"name": "sample-item",
		"request": {
			"method": "GET",
			"url": "https://some-url.com"
		},
		"variable": [{}]
	}
]
`

func TestITem_invalids(t *testing.T) {
	var i []Item
	testutil.Decode(invalidItemsString, &i, t)
	for _, invalidItem := range i {
		testutil.FatalIfNoError(invalidItem.InitAndValidate(), t)
	}
}

func TestItem_isEmpty(t *testing.T) {
	i := Item{}
	testutil.FatalIfFalse("Item is empty", i.IsEmpty(), t)
}
