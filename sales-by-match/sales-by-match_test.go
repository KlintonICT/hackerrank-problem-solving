package sales_by_match

import (
	"testing"
)

func TestSockMerchant(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := sockMerchant(tc.data.n, tc.data.ar)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
