package jumping_on_the_clouds

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	jumps, index := int32(0), 0

	for index < len(c) {
		step1 := index + 1
		step2 := index + 2

		if step2 < len(c) && c[step2] == 0 {
			jumps += 1
			index = step2
		} else if step1 < len(c) && c[step1] == 0 {
			jumps += 1
			index = step1
		} else {
			break
		}
	}

	return jumps
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Jumping on the Clouds")

	for _, tc := range tc.getTestCases() {
		result := jumpingOnClouds(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
