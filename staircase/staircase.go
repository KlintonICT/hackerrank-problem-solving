package staircase

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	for i := int32(1); i <= n; i++ {
		fmt.Println(strings.Repeat(" ", int(n-i)) + strings.Repeat("#", int(i)))
	}
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Staircase")

	for _, tc := range tc.getTestCases() {
		fmt.Printf("TestCase: %d; data: %d; result: \n", tc.id, tc.data)
		staircase(tc.data)
	}
}
