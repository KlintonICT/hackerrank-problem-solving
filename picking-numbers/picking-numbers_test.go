package picking_numbers

import (
	"testing"
)

func TestPickingNumbers(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := pickingNumbers(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
