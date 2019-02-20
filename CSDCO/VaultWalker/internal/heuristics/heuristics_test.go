package heuristics

import (
	"testing"
)

func Test(t *testing.T) {
	ht := CSDCOHTs()
	if len(ht) < 1 {
		t.Error("Expected an array of test of 1 or larger")
	}
}
