package library_fine

import (
	"testing"
)

func TestLibraryFine(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := libraryFine(tc.data.d1, tc.data.m1, tc.data.y1, tc.data.d2, tc.data.m2, tc.data.y2)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
