package compare_the_triplets

import (
	"reflect"
	"testing"
)

func TestCompareTriplets(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := compareTriplets(tc.data.a, tc.data.b)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
