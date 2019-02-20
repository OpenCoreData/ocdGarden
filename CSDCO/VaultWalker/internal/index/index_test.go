package index

import (
	"testing"
)

func Test(t *testing.T) {
	if caselessContains("THIS", "this") != true {
		t.Error("Expected THIS to equal this")
	}
}
