package divisible_sum_pairs

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	count := int32(0)

	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				count += 1
			}
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Divisible Sum Pairs")

	for _, tc := range tc.getTestCases() {
		result := divisibleSumPairs(tc.data.n, tc.data.k, tc.data.ar)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
