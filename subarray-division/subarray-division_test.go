package subarray_division

import "testing"

func TestBirthday(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := birthday(tc.data.s, tc.data.d, tc.data.m)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
