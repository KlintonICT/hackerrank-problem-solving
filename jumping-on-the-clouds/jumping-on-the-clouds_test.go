package jumping_on_the_clouds

import (
	"reflect"
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := jumpingOnClouds(tc.data)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
