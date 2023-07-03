package mars_exploration

import (
	"testing"
)

func TestMarsExploration(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := marsExploration(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
