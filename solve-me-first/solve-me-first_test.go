package solve_me_first

import (
	"testing"
)

func TestWeightedUniformStrings(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := solveMeFirst(tc.data.a, tc.data.b)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
