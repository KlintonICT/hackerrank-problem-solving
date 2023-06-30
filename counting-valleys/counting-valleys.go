package counting_valleys

import "fmt"

func countingValleys(steps int32, path string) int32 {
	isValley, valleyCount, count := false, 0, 0

	for _, char := range path {
		if char == 'U' {
			count++
		} else {
			count--
		}

		if count < 0 {
			isValley = true
		} else if count == 0 && isValley {
			valleyCount++
			isValley = false
		}
	}

	return int32(valleyCount)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Counting Valleys")

	for _, tc := range tc.getTestCases() {
		result := countingValleys(tc.data.steps, tc.data.path)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
