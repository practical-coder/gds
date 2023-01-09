package words

import "testing"

func TestAlphabetize(t *testing.T) {
	given := "pirates"
	expected := "aeiprst"
	result := alphabetize(given)
	if result != expected {
		t.Errorf("given: %s, expected: %s", given, expected)
	}
}
