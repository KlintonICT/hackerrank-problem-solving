package sherlock_and_squares

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
	start := int32(math.Ceil(math.Sqrt(float64(a))))
	end := int32(math.Floor(math.Sqrt(float64(b))))

	return end - start + 1
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Sherlock and Squares")

	for _, tc := range tc.getTestCases() {
		result := squares(tc.data.a, tc.data.b)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
