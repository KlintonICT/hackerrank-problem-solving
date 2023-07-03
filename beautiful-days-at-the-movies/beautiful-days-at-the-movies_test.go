package beautiful_days_at_the_movies

import (
	"testing"
)

func TestBeautifulDays(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := beautifulDays(tc.data.i, tc.data.j, tc.data.k)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
