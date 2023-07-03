package pangrams

import (
	"testing"
)

func TestPangrams(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := pangrams(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
