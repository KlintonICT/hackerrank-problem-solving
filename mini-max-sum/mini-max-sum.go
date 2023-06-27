package mini_max_sum

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	min, max := 0, 0

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr)-1; i++ {

		min += int(arr[i])
		max += int(arr[i+1])
	}

	fmt.Printf("%d %d\n", min, max)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Staircase")

	for _, tc := range tc.getTestCases() {
		fmt.Printf("TestCase: %d; data: %d; result: \n", tc.id, tc.data)
		miniMaxSum(tc.data)
	}
}
