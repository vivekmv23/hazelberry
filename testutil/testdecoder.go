package testutil

import (
	"encoding/json"
	"strings"
	"testing"
)

func Decode(data string, targetPointer any, t *testing.T) {
	reader := strings.NewReader(data)
	decoder := json.NewDecoder(reader)
	FatalIfError(decoder.Decode(targetPointer), t)
}
