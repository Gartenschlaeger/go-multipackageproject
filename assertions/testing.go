package assertions

import "testing"

func AssertEqual(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected \"%v\" but got \"%v\".\n", expected, actual)
	}
}
