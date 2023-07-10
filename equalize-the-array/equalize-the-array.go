package equalize_the_array

import "fmt"

func equalizeArray(arr []int32) int32 {
	pairs := make(map[int32]int32)
	max := int32(0)

	for _, element := range arr {
		pairs[element] += 1
		if pairs[element] > max {
			max = pairs[element]
		}
	}

	return int32(len(arr)) - max
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Equalize the Array")

	for _, tc := range tc.getTestCases() {
		result := equalizeArray(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
