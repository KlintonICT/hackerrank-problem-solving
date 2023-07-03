package hackerrank_in_a_string

import "fmt"

func hackerrankInString(s string) string {
	word := "hackerrank"

	for _, char := range s {
		if char == []rune(word)[0] {
			word = word[1:]
		}

		if len(word) == 0 {
			return "YES"
		}
	}

	return "NO"
}

func Run() {
	tc := &TestCase{}

	fmt.Println("HackerRank in a String!")

	for _, tc := range tc.getTestCases() {
		result := hackerrankInString(tc.data)

		fmt.Printf("TestCase: %d; data: %s; result: %s\n", tc.id, tc.data, result)
	}
}
