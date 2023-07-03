package cats_and_a_mouse

import (
	"testing"
)

func TestCatAndMouse(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := catAndMouse(tc.data.x, tc.data.y, tc.data.z)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
