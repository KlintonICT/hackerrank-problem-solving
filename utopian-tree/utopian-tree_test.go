package utopian_tree

import (
	"testing"
)

func TestUtopianTree(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := utopianTree(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
