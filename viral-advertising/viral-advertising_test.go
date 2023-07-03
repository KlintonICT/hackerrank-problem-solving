package viral_advertising

import (
	"testing"
)

func TestAngryProfessor(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := viralAdvertising(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
