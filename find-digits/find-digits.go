package find_digits

import (
	"fmt"
	"strconv"
	"strings"
)

func findDigits(n int32) int32 {
	str := strconv.Itoa(int(n))
	arr := strings.Split(str, "")
	count := int32(0)

	for _, element := range arr {
		divisor, _ := strconv.Atoi(element)

		if divisor != 0 && n%int32(divisor) == 0 {
			count += 1
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Find Digits")

	for _, tc := range tc.getTestCases() {
		result := findDigits(tc.data)

		fmt.Printf("TestCase: %d; data: %d; result: %d\n", tc.id, tc.data, result)
	}
}
