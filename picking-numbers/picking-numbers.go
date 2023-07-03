package picking_numbers

import (
	"fmt"
	"math"
	"sort"
)

func pickingNumbers(a []int32) int32 {
	longestSubArr := 0

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	for i := 0; i < len(a); i++ {
		var subArr []int32
		for j := i + 1; j < len(a); j++ {

			if math.Abs(float64(a[i]-a[j])) <= 1 {
				subArr = append(subArr, a[j])
			}
		}

		if longestSubArr < len(subArr) {
			longestSubArr = len(subArr)
		}
	}

	return int32(longestSubArr + 1)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Picking Numbers")

	for _, tc := range tc.getTestCases() {
		result := pickingNumbers(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
