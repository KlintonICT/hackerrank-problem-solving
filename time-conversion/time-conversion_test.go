package time_conversion

import "testing"

func TestTimeConversion(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := timeConversion(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, result, tc.expected)
		}
	}
}
