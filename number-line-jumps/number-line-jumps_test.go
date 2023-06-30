package number_line_jumps

import (
	"testing"
)

func TestKangaroo(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := kangaroo(tc.data.x1, tc.data.v1, tc.data.x2, tc.data.v2)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
