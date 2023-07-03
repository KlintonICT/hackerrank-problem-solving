package into_to_tutorial_challenges

import "fmt"

func introTutorial(V int32, arr []int32) int32 {
	for i, item := range arr {
		if item == V {
			return int32(i)
		}
	}

	return -1
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Intro to Tutorial Challenges")

	for _, tc := range tc.getTestCases() {
		result := introTutorial(tc.data.V, tc.data.arr)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
