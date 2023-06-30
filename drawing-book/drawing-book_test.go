package drawing_book

import (
	"testing"
)

func TestPageCount(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := pageCount(tc.data.n, tc.data.p)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
