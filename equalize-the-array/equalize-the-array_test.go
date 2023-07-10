package equalize_the_array

import (
	"testing"
)

func TestEqualizeArray(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := equalizeArray(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
