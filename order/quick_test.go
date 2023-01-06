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
	Quick[string](names)
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
