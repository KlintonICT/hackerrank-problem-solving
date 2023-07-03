package electronics_shop

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	result := int32(-1)

	for _, k := range keyboards {
		for _, d := range drives {
			sum := k + d

			if sum <= b && sum > result {
				result = sum
			}
		}
	}

	return result
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Electronics Shop")

	for _, tc := range tc.getTestCases() {
		result := getMoneySpent(tc.data.keyboards, tc.data.drives, tc.data.b)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
