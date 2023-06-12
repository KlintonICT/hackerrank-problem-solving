package simple_array_sum

import "fmt"

func simpleArraySum(ar []int32) int32 {
	sum := int32(0)

	for _, element := range ar {
		sum += element
	}

	return sum
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Simple Array Sum")

	for _, tc := range tc.getTestCases() {
		result := simpleArraySum(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
