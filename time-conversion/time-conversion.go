package time_conversion

import (
	"fmt"
	"time"
)

func timeConversion(s string) string {
	t, _ := time.Parse("03:04:05PM", s)

	return t.Format("15:04:05")
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Time Conversion")

	for _, tc := range tc.getTestCases() {
		result := timeConversion(tc.data)

		fmt.Printf("TestCase: %d; data: %s; result: %s\n", tc.id, tc.data, result)
	}
}
