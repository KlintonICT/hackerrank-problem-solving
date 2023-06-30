package subarray_division

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	if m == 1 && len(s) == 1 && s[0] == d {
		return 1
	}

	if m > 1 && len(s) > 1 {
		count := int32(0)

		for i := 0; i <= len(s)-int(m); i++ {
			tempCount := s[i]

			for j := i + 1; j < i+int(m); j++ {
				tempCount += s[j]
			}

			if tempCount == d {
				count += 1
			}
		}

		return count
	}

	return 0
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Subarray division")

	for _, tc := range tc.getTestCases() {
		result := birthday(tc.data.s, tc.data.d, tc.data.m)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
