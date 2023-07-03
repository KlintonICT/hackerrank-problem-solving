package find_digits

import (
	"testing"
)

func TestFindDigits(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := findDigits(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
