package minimum_absolute_difference

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	mindDiff := int32(math.MaxInt32)

	for i := 1; i < len(arr); i++ {
		diff := int32(math.Abs(float64(arr[i] - arr[i-1])))

		if diff < mindDiff {
			mindDiff = diff
		}
	}

	return mindDiff
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Minimum Absolute Difference in an Array")

	for _, tc := range tc.getTestCases() {
		result := minimumAbsoluteDifference(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
