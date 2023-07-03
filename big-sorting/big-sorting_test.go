package big_sorting

import (
	"reflect"
	"testing"
)

func TestBigSorting(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := bigSorting(tc.data)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
