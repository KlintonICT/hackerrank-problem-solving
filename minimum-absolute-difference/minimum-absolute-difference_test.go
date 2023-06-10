package minimum_absolute_difference

import (
	"testing"
)

func TestMinimumAbsoluteDifference(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := minimumAbsoluteDifference(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
