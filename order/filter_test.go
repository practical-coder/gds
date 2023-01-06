package order

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilter(t *testing.T) {
	floats := []float64{14.1, 15.7, 16.3, 15.8, 83.2}

	filterFunc := func(item float64) bool {
		return item < 16
	}
	filtered := Filter(floats, filterFunc)
	expected := []float64{14.1, 15.7, 15.8}
	if !cmp.Equal(filtered, expected) {
		t.Errorf("filtered: %v\nexpected: %v\n", filtered, expected)
	}

}
