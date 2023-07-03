package sequence_equation

import "fmt"

func findIndex(arr []int32, target int) int {
	for i, element := range arr {
		if element == int32(target) {
			return i
		}
	}

	return -1
}

func permutationEquation(p []int32) []int32 {
	tempArr := make([]int32, len(p))
	result := make([]int32, len(p))

	for i := 0; i < len(p); i++ {
		target := findIndex(p, i+1)
		tempArr[i] = int32(target) + 1
	}

	for i := 0; i < len(tempArr); i++ {
		target := int32(findIndex(p, int(tempArr[i])))
		result[i] = target + 1
	}

	return result
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Sequence Equation")

	for _, tc := range tc.getTestCases() {
		result := permutationEquation(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %v\n", tc.id, tc.data, result)
	}
}
