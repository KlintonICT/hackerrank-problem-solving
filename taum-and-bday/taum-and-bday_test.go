package taum_and_bday

import (
	"testing"
)

func TestTaumBday(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := taumBday(tc.data.b, tc.data.w, tc.data.bc, tc.data.wc, tc.data.z)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
