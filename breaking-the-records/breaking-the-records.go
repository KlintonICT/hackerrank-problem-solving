package breaking_the_records

import "fmt"

func breakingRecords(scores []int32) []int32 {
	min, max, count := scores[0], scores[0], []int32{0, 0}

	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			count[0] += 1
		}
		if scores[i] < min {
			min = scores[i]
			count[1] += 1
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Breaking the records")

	for _, tc := range tc.getTestCases() {
		result := breakingRecords(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %v\n", tc.id, tc.data, result)
	}
}
