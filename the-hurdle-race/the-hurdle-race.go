package the_hurdle_race

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
	maxHeight := int32(0)

	for _, h := range height {
		if maxHeight < h {
			maxHeight = h
		}
	}

	if maxHeight > k {
		return maxHeight - k
	}

	return 0
}

func Run() {
	tc := &TestCase{}

	fmt.Println("A Very Big Sum")

	for _, tc := range tc.getTestCases() {
		result := hurdleRace(tc.data.k, tc.data.height)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
