package compare_the_triplets

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	result := []int32{0, 0}

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			result[0] += 1
		} else if a[i] < b[i] {
			result[1] += 1
		}
	}

	return result
}

func Run() {
	fmt.Println("Compare the Triplets")
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		result := compareTriplets(tc.data.a, tc.data.b)

		fmt.Printf("TestCase: %d; data: %+v; result: %v\n", tc.id, tc.data, result)
	}
}
