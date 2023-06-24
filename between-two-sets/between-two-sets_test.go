package between_two_sets

import (
	"testing"
)

func TestGetTotalX(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := getTotalX(tc.data.a, tc.data.b)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
