package acm_icpc_team

import (
	"reflect"
	"testing"
)

func TestAcmTeam(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := acmTeam(tc.data)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TestCase: %d; actual: %v; expected: %v", tc.id, result, tc.expected)
		}
	}
}
