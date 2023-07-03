package sequence_equation

import (
	"reflect"
	"testing"
)

func TestPermutationEquation(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := permutationEquation(tc.data)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
