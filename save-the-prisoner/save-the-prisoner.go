package save_the_prisoner

import "fmt"

func saveThePrisoner(n int32, m int32, s int32) int32 {
	return int32((m+s-2)%n + 1)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Save the Prisoner!")

	for _, tc := range tc.getTestCases() {
		result := saveThePrisoner(tc.data.n, tc.data.m, tc.data.s)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
