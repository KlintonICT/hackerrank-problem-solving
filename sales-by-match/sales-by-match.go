package sales_by_match

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	pairs, count := []int32{ar[0]}, int32(0)

	for i := int32(1); i < n; i++ {
		found := false

		for j := 0; j < len(pairs); j++ {
			if ar[i] == pairs[j] {
				pairs = append(pairs[:j], pairs[j+1:]...)
				count += 1
				found = true
				break
			}
		}

		if !found {
			pairs = append(pairs, ar[i])
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Sales by Match")

	for _, tc := range tc.getTestCases() {
		result := sockMerchant(tc.data.n, tc.data.ar)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
