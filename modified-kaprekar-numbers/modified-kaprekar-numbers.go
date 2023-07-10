package modified_kaprekar_numbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func kaprekarNumbers(p int32, q int32) {
	result := []string{}

	for n := p; n <= q; n++ {
		square := strconv.Itoa(int(math.Pow(float64(n), 2)))
		left, _ := strconv.Atoi(square[0:int(math.Floor(float64(len(square)/2)))])
		right, _ := strconv.Atoi(square[int(math.Floor(float64(len(square)/2))):])

		if left+right == int(n) {
			result = append(result, strconv.Itoa(int(n)))
		}
	}

	if len(result) > 0 {
		fmt.Println(strings.Join(result, " "))
	} else {
		fmt.Println("INVALID RANGE")
	}
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Modified Kaprekar Numbers")

	for _, tc := range tc.getTestCases() {

		fmt.Printf("TestCase: %d; data: %+v; result: \n", tc.id, tc.data)
		kaprekarNumbers(tc.data.p, tc.data.q)
	}
}
