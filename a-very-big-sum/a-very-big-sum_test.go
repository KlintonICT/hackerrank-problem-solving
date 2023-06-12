package a_very_big_sum

import (
	"testing"
)

func TestAngryProfessor(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := aVeryBigSum(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
