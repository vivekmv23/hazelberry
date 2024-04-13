package utils

import (
	"testing"
)

func FatalIfError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func FatalIfNoError(err error, t *testing.T) {
	if err == nil {
		t.Fatal("Expected some error but got none")
	}
}

func FatalIfTrue(key string, val bool, t *testing.T) {
	if val {
		t.Fatalf("Expected %s to be true, but is false\n", key)
	}
}

func FatalIfFalse(key string, val bool, t *testing.T) {
	if !val {
		t.Fatalf("Expected %s to be false, but is true\n", key)
	}
}

func FatalIfNotEquals(expected string, result string, t *testing.T) {
	if expected != result {
		t.Fatalf("Expected \"%s\" but got \"%s\"\n", expected, result)
	}
}
