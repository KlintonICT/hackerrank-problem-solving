package apple_and_orange

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	numOfApple, numOfOrange := 0, 0

	for _, apple := range apples {
		distance := a + apple
		if distance >= s && distance <= t {
			numOfApple += 1
		}
	}

	for _, orange := range oranges {
		distance := b + orange
		if distance >= s && distance <= t {
			numOfOrange += 1
		}
	}

	fmt.Printf("%d\n%d\n", numOfApple, numOfOrange)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Apple and Orange")

	for _, tc := range tc.getTestCases() {
		fmt.Printf("TestCase: %d; data: %+v; result: \n", tc.id, tc.data)
		countApplesAndOranges(tc.data.s, tc.data.t, tc.data.a, tc.data.b, tc.data.apples, tc.data.oranges)
	}
}
