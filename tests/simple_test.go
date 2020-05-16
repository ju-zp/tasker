package tests

import (
	"fmt"
	"testing"
)

func Test_simple(t *testing.T) {
	fmt.Print("Success")
	if false {
		t.Fatal("Expected true to not be false")
	}
	t.Log("True is true")
}
