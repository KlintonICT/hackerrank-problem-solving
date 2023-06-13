package diagonal_difference

import "testing"

func TestDiagonalDifference(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := diagonalDifference(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %v; expected: %d", tc.id, result, tc.expected)
		}
	}
}
