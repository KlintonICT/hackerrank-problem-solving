package cut_the_sticks

import (
	"reflect"
	"testing"
)

func TestCutTheSticks(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := cutTheSticks(tc.data)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
