package breaking_the_records

import (
	"reflect"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := breakingRecords(tc.data)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
