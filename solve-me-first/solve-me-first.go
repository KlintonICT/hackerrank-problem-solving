package solve_me_first

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Solve Me First")

	for _, tc := range tc.getTestCases() {
		result := solveMeFirst(tc.data.a, tc.data.b)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
