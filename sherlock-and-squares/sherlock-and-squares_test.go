package sherlock_and_squares

import (
	"testing"
)

func TestSquares(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := squares(tc.data.a, tc.data.b)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
