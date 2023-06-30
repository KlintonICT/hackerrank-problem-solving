package divisible_sum_pairs

import (
	"testing"
)

func TestDivisibleSumPairs(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := divisibleSumPairs(tc.data.n, tc.data.k, tc.data.ar)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
