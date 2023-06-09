package circular_array_rotation

import (
	"reflect"
	"testing"
)

func testCircularArrayRotationSolution(t *testing.T, solutionFunc SolutionFuncType) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := solutionFunc(tc.data.a, tc.data.k, tc.data.queries)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}

func TestCircularArrayRotationSolution1(t *testing.T) {
	testCircularArrayRotationSolution(t, circularArrayRotationSolution1)
}

func TestCircularArrayRotationSolution2(t *testing.T) {
	testCircularArrayRotationSolution(t, circularArrayRotationSolution2)
}
