package simple_array_sum

import "testing"

func TestWeightedUniformStrings(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := simpleArraySum(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
