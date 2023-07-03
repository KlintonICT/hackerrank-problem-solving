package electronics_shop

import (
	"testing"
)

func TestGetMoneySpent(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := getMoneySpent(tc.data.keyboards, tc.data.drives, tc.data.b)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
