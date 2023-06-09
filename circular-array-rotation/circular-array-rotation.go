package circular_array_rotation

import "fmt"

type SolutionFuncType func([]int32, int32, []int32) []int32

func circularArrayRotationSolution1(a []int32, k int32, queries []int32) []int32 {
	rotate := a
	var result []int32

	for i := int32(1); i <= k; i++ {
		rotate = append(rotate[len(a)-1:], rotate[0:len(a)-1]...)
	}

	for _, query := range queries {
		result = append(result, rotate[query])
	}

	return result
}

func circularArrayRotationSolution2(a []int32, k int32, queries []int32) []int32 {
	var result []int32
	lenA := int32(len(a))

	for _, query := range queries {
		index := query - (k % lenA)

		if index < 0 {
			index += lenA
		}

		result = append(result, a[index])
	}

	return result
}

func runSolution(solutionFunc SolutionFuncType) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := solutionFunc(tc.data.a, tc.data.k, tc.data.queries)

		fmt.Printf("TestCase: %d; data: %+v; result: %v\n", tc.id, tc.data, result)
	}
}

func Run() {
	fmt.Println("Circular Array Rotation")

	fmt.Println("Solution 1")
	runSolution(circularArrayRotationSolution1)

	fmt.Println("Solution 2")
	runSolution(circularArrayRotationSolution2)
}
