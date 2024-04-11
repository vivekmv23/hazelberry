package utils

import (
	"fmt"
	"testing"
)

func Equals(expected string, result string) (msg string) {
	if expected != result {
		return fmt.Sprintf("Expected \"%s\" but got \"%s\"\n", expected, result)
	}
	return ""
}

func NoError(err error) (msg string) {
	if err != nil {
		return fmt.Sprintf("Expected no error but got: %v", err)
	}
	return ""
}

func HasError(err error) (msg string) {
	if err == nil {
		return "Expected some error but got none"
	}
	return ""
}

func IsTrue(key string, val bool) (msg string) {
	if !val {
		return fmt.Sprintf("Expected %s to be true, but is false\n", key)
	}
	return ""
}

func IsFalse(key string, val bool) (msg string) {
	if val {
		return fmt.Sprintf("Expected %s to be false, but is true\n", key)
	}
	return ""
}

func FatalIfMsg(msg string, t *testing.T) {
	if msg != "" {
		t.Fatal(msg)
	}
}
