package jumping_on_the_clouds_revisited

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	position := k % int32(len(c))
	energy := 100 - (1 + c[position]*2)

	for position != 0 {
		position = (position + k) % int32(len(c))
		energy -= 1 + c[position]*2
	}

	return energy
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Jumping on the Clouds: Revisited")

	for _, tc := range tc.getTestCases() {
		result := jumpingOnClouds(tc.data.c, tc.data.k)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
