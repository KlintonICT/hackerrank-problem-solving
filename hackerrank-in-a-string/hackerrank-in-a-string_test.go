package hackerrank_in_a_string

import (
	"testing"
)

func TestHackerrankInString(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := hackerrankInString(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
