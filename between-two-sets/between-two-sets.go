package between_two_sets

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	min := a[len(a)-1]
	max := b[0]
	count := int32(0)

	for i := min; i <= max; i++ {
		isAllFactorA, isAllFactorB := true, true

		for _, numA := range a {
			if i%numA != 0 {
				isAllFactorA = false
				break
			}
		}

		for _, numB := range b {
			if numB%i != 0 {
				isAllFactorB = false
				break
			}
		}

		if isAllFactorA && isAllFactorB {
			count++
		}
	}

	return count
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Between Two Sets")

	for _, tc := range tc.getTestCases() {
		result := getTotalX(tc.data.a, tc.data.b)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
