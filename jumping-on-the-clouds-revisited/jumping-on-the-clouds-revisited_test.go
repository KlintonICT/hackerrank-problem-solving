package jumping_on_the_clouds_revisited

import (
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := jumpingOnClouds(tc.data.c, tc.data.k)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
