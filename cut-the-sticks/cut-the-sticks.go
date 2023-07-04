package cut_the_sticks

import (
	"fmt"
	"sort"
)

func cutTheSticks(arr []int32) []int32 {
	sortArr, result := make([]int32, len(arr)), []int32{}
	copy(sortArr, arr)

	sort.Slice(sortArr, func(i, j int) bool {
		return sortArr[i] < sortArr[j]
	})

	tempSortArr := []int32{}
	result = append(result, int32(len(sortArr)))

	for i := 0; i < len(arr); i++ {
		for _, item := range sortArr {
			remain := item - sortArr[0]
			if remain > 0 {
				tempSortArr = append(tempSortArr, remain)
			}
		}

		if len(tempSortArr) == 0 {
			break
		} else {
			result = append(result, int32(len(tempSortArr)))
			sortArr = tempSortArr
			tempSortArr = []int32{}
		}
	}

	return result
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Cut the Sticks")

	for _, tc := range tc.getTestCases() {
		result := cutTheSticks(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %v\n", tc.id, tc.data, result)
	}
}
