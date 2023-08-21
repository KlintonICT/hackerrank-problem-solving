package beautiful_triplets

import "fmt"

func isExist(target int32, arr []int32) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}

	return false
}

func beautifulTriplets(d int32, arr []int32) int32 {
	count := int32(0)

	for _, element := range arr {
		j := element + d
		k := j + d

		if isExist(j, arr) && isExist(k, arr) {
			count += 1
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Beautiful Triplets")

	for _, tc := range tc.getTestCases() {
		result := beautifulTriplets(tc.data.d, tc.data.arr)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
