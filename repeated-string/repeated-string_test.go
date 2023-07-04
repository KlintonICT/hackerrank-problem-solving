package repeated_string

import (
	"testing"
)

func TestRepeatedString(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := repeatedString(tc.data.s, tc.data.n)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
