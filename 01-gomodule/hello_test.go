package hello

import (
	"testing"
)

func TestMain(t *testing.T) {
	expected := 54
	actual := 54
	if expected != actual {
		t.Errorf("value not expected %v", expected)
	}
}
