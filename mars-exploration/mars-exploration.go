package mars_exploration

import "fmt"

func marsExploration(s string) int32 {
	alteredLetter := 0

	for i, ch := range s {
		if ((i%3 == 0 || i%3 == 2) && ch != 'S') || (i%3 == 1 && ch != 'O') {
			alteredLetter += 1
		}
	}

	return int32(alteredLetter)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Mars Exploration")

	for _, tc := range tc.getTestCases() {
		result := marsExploration(tc.data)

		fmt.Printf("TestCase: %d; data: %s; result: %d\n", tc.id, tc.data, result)
	}
}
