package pangrams

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
	sentence := strings.ToLower(s)

	alphabetMap := make(map[rune]bool)

	for _, char := range sentence {
		if char >= 'a' && char <= 'z' {
			alphabetMap[char] = true
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		if !alphabetMap[i] {
			return "not pangram"
		}
	}

	return "pangram"
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Pangrams")

	for _, tc := range tc.getTestCases() {
		result := pangrams(tc.data)

		fmt.Printf("TestCase: %d; data: %s; result: %s\n", tc.id, tc.data, result)
	}
}
