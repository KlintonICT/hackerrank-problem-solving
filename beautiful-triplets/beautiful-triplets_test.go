package beautiful_triplets

import (
	"testing"
)

func TestBeautifulTriplets(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := beautifulTriplets(tc.data.d, tc.data.arr)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
