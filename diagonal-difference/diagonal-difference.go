package diagonal_difference

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	leftSum, rightSum := int32(0), int32(0)
	arrLen := len(arr[0])

	for i := 0; i < arrLen; i++ {
		leftSum += arr[i][i]
		rightSum += arr[i][arrLen-i-1]
	}

	return int32(math.Abs(float64(leftSum - rightSum)))
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Diagonal Difference")

	for _, tc := range tc.getTestCases() {
		result := diagonalDifference(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
