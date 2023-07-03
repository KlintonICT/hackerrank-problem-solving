package the_hurdle_race

import (
	"testing"
)

func TestHurdleRace(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := hurdleRace(tc.data.k, tc.data.height)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
