package angry_professor

import (
	"testing"
)

func TestAngryProfessor(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := angryProfessor(tc.data.k, tc.data.a)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
