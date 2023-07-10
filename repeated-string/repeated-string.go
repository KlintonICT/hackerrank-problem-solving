package repeated_string

import "fmt"

func repeatedString(s string, n int64) int64 {
	occurrence := int64(0)
	remainder := n % int64(len(s))
	quotient := n / int64(len(s))

	for _, char := range s {
		if char == 'a' {
			occurrence += 1
		}
	}

	occurrence = occurrence * int64(quotient)

	if remainder > 0 {
		for i := int64(0); i < remainder; i++ {
			if s[i] == 'a' {
				occurrence += 1
			}
		}
	}

	return occurrence
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Repeated String")

	for _, tc := range tc.getTestCases() {
		result := repeatedString(tc.data.s, tc.data.n)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
