package utils_test

import (
	"testing"
)

func TestFoo(t *testing.T) {
	defer func() {
		t.Log("defer")
	}()

	// t.Log("2")
	// t.FailNow()

	t.Log("1")
	t.Fatal("something happen")
	t.Log("2")
}
