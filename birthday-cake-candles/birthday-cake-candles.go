package birthday_cake_candles

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	count, max := int32(1), candles[0]

	for i := 1; i < len(candles); i++ {
		if candles[i] > max {
			max = candles[i]
			count = 1
		} else if candles[i] == max {
			count++
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Birthday Cake Candles")

	for _, tc := range tc.getTestCases() {
		result := birthdayCakeCandles(tc.data)

		fmt.Printf("TestCase: %d; data: %d; result: %d\n", tc.id, tc.data, result)
	}
}
