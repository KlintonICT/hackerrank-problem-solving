package grading_students

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := gradingStudents(tc.data)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
