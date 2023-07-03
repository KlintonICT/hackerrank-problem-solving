package big_sorting

import (
	"fmt"
	"sort"
)

func bigSorting(unsorted []string) []string {
	arr := make([]string, len(unsorted))
	copy(arr, unsorted)

	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i]) == len(arr[j]) {
			return arr[i] < arr[j]
		}

		return len(arr[i]) < len(arr[j])
	})

	return arr
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Big Sorting")

	for _, tc := range tc.getTestCases() {
		result := bigSorting(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %v\n", tc.id, tc.data, result)
	}
}
