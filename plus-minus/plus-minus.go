package plus_minus

import (
	"fmt"
)

func plusMinus(arr []int32) {
	neg, pos, zero, len := 0, 0, 0, len(arr)

	for _, element := range arr {
		if element > 0 {
			pos += 1
		} else if element < 0 {
			neg += 1
		} else {
			zero += 1
		}
	}

	fmt.Printf("%.6f\n", float64(pos)/float64(len))
	fmt.Printf("%.6f\n", float64(neg)/float64(len))
	fmt.Printf("%.6f\n", float64(zero)/float64(len))
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Plus Minus")

	for _, tc := range tc.getTestCases() {
		fmt.Printf("TestCase: %d; data: %v; result: \n", tc.id, tc.data)
		plusMinus(tc.data)
	}
}
