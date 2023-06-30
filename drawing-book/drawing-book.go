package drawing_book

import (
	"fmt"
	"math"
)

func pageCount(n int32, p int32) int32 {
	numOfPages := n

	if n%2 == 0 {
		numOfPages += 1
	}

	fromFront := math.Floor(float64(p / 2))
	fromBack := math.Floor(float64(numOfPages-p) / 2)

	if fromFront < fromBack {
		return int32(fromFront)
	}

	return int32(fromBack)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Drawing Book")

	for _, tc := range tc.getTestCases() {
		result := pageCount(tc.data.n, tc.data.p)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
