package birthday_cake_candles

import (
	"testing"
)

func TestBirthdayCakeCandles(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := birthdayCakeCandles(tc.data)

		if result != tc.expected {
			t.Errorf("TestCase: %d; actual: %d; expected: %d", tc.id, result, tc.expected)
		}
	}
}
