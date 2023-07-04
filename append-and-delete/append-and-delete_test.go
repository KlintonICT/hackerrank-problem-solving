package append_and_delete

import (
	"testing"
)

func TestAppendAndDelete(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := appendAndDelete(tc.data.s, tc.data.t, tc.data.k)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
