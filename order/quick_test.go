package order

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuicksort(t *testing.T) {
	names := []string{
		"Zuza",
		"Asia",
		"Basia",
		"Kasia",
		"Magda",
		"Iza",
	}
	Quicksort[string](names, 0, len(names)-1)
	expected := []string{
		"Asia",
		"Basia",
		"Iza",
		"Kasia",
		"Magda",
		"Zuza",
	}
	if !cmp.Equal(names, expected) {
		t.Errorf("names: %v\nexpected: %v\n", names, expected)
	}
}
