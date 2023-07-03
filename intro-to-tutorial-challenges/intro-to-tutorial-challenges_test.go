package into_to_tutorial_challenges

import (
	"testing"
)

func TestIntroTutorial(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := introTutorial(tc.data.V, tc.data.arr)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
