package forming_a_magic_square

import (
	"fmt"
	"math"
)

func formingMagicSquare(s [][]int32) int32 {
	magicSquares := []([][]int32){
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}
	cost := int32(math.MaxInt32)

	for _, magicSquare := range magicSquares {
		min := int32(0)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				min += int32(math.Abs(float64(s[i][j] - magicSquare[i][j])))
			}
		}

		if min == 0 {
			cost = 0
			break
		} else {
			if cost > min {
				cost = min
			}
		}
	}

	return cost
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Forming a Magic Square")

	for _, tc := range tc.getTestCases() {
		result := formingMagicSquare(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
