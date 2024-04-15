package utils

import (
	"fmt"
	"testing"
)

var VERBOSE_TEST_LOG bool = false

func FatalIfError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
	printIfVerbose("No error receieved as expected")
}

func FatalIfNoError(err error, t *testing.T) {
	if err == nil {
		t.Fatal("Expected some error but got none")
	}
	printIfVerbose("Error recevied as expected:", err)
}

func FatalIfTrue(key string, val bool, t *testing.T) {
	if val {
		t.Fatalf("Expected %s to be false, but is true\n", key)
	}
	printIfVerbose(fmt.Sprintf("%s is false as expected\n", key))
}

func FatalIfFalse(key string, val bool, t *testing.T) {
	if !val {
		t.Fatalf("Expected %s to be true, but is false\n", key)
	}
	printIfVerbose(fmt.Sprintf("%s is true as expected\n", key))
}

func FatalIfNotEquals(expected string, result string, t *testing.T) {
	if expected != result {
		t.Fatalf("Expected \"%s\" but got \"%s\"\n", expected, result)
	}
	printIfVerbose("Both expected and receieved string are equal:", result)
}

func printIfVerbose(a ...any) {
	if VERBOSE_TEST_LOG {
		fmt.Println(a...)
	}
}
