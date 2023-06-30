package counting_valleys

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := countingValleys(tc.data.steps, tc.data.path)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
