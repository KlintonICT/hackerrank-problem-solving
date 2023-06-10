package weighted_uniform_strings

import (
	"reflect"
	"testing"
)

func TestWeightedUniformStrings(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := weightedUniformStrings(tc.data.s, tc.data.queries)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
