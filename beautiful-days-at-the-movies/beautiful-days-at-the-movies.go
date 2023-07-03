package beautiful_days_at_the_movies

import (
	"fmt"
	"strconv"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	count := int32(0)

	for num := i; num <= j; num++ {
		numStr := strconv.Itoa(int(num))
		reversedStr := ""

		for index := len(numStr) - 1; index >= 0; index-- {
			reversedStr += string(numStr[index])
		}

		reversedNum, _ := strconv.Atoi(reversedStr)

		if ((num - int32(reversedNum)) % k) == 0 {
			count++
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Beautiful Days at the Movies")

	for _, tc := range tc.getTestCases() {
		result := beautifulDays(tc.data.i, tc.data.j, tc.data.k)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
