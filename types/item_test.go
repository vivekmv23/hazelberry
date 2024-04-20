package types

import (
	"encoding/json"
	"strings"
	"testing"

	testingutils "github.com/vivekmv23/hazelberry/utils"
)

var validItemString string = `
{
	"name": "sample-item",
	"id": "sample-id",
	"request": {
		"method": "GET",
		"url": "https://some-url.some-domain.com?some-param-key=some-param-value"
	}
}
`

func TestItem_valid(t *testing.T) {
	reader := strings.NewReader(validItemString)
	decoder := json.NewDecoder(reader)
	var i Item
	testingutils.FatalIfError(decoder.Decode(&i), t)
	err := i.InitAndValidate()
	testingutils.FatalIfError(err, t)
	testingutils.FatalIfTrue("Item is empty", i.IsEmpty(), t)
}

var invalidItemsString string = `
[
	{
		"request": {
			"method": "GET",
			"url": "https://some-url.some-domain.com?some-param-key=some-param-value"
		}
	},
	{
		"name": "sample-item"
	}
]
`

func TestITem_invalids(t *testing.T) {
	reader := strings.NewReader(invalidItemsString)
	decoder := json.NewDecoder(reader)
	var i []Item
	testingutils.FatalIfError(decoder.Decode(&i), t)
	for _, invalidItem := range i {
		testingutils.FatalIfNoError(invalidItem.InitAndValidate(), t)
	}
}

func TestItem_isEmpty(t *testing.T) {
	i := Item{}
	testingutils.FatalIfFalse("Item is empty", i.IsEmpty(), t)
}
