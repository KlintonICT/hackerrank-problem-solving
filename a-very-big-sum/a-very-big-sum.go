package a_very_big_sum

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	result := int64(0)

	for _, element := range ar {
		result += element
	}

	return result
}

func Run() {
	tc := &TestCase{}

	fmt.Println("A Very Big Sum")

	for _, tc := range tc.getTestCases() {
		result := aVeryBigSum(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
