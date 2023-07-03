package viral_advertising

import "fmt"

func viralAdvertising(n int32) int32 {
	liked := 2
	cumulative := liked

	for i := int32(2); i <= n; i++ {
		liked = (liked * 3) / 2
		cumulative += liked
	}

	return int32(cumulative)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Viral Advertising")

	for _, tc := range tc.getTestCases() {
		result := viralAdvertising(tc.data)

		fmt.Printf("TestCase: %d; data: %d; result: %d\n", tc.id, tc.data, result)
	}
}
