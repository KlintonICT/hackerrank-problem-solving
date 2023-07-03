package utopian_tree

import "fmt"

func utopianTree(n int32) int32 {
	grownTree := int32(1)

	for i := int32(1); i <= n; i++ {
		if i%2 == 0 {
			grownTree += 1
		} else {
			grownTree *= 2
		}
	}

	return grownTree
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Utopian Tree")

	for _, tc := range tc.getTestCases() {
		result := utopianTree(tc.data)

		fmt.Printf("TestCase: %d; data: %d; result: %d\n", tc.id, tc.data, result)
	}
}
