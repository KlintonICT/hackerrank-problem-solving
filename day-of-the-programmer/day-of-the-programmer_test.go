package day_of_the_programmer

import (
	"testing"
)

func TestAngryProfessor(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := dayOfProgrammer(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
