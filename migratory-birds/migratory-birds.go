package migratory_birds

import "fmt"

func migratoryBirds(arr []int32) int32 {
	obj, max, smallType := make(map[int32]int32), int32(0), int32(0)

	for _, element := range arr {
		if _, ok := obj[element]; ok {
			obj[element] += 1

			if obj[element] > max {
				max = obj[element]
				smallType = element
			}
		} else {
			obj[element] = 1
		}
	}

	for key, value := range obj {
		if key < smallType && value == max {
			smallType = key
		}
	}

	return smallType
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Migratory Birds")

	for _, tc := range tc.getTestCases() {
		result := migratoryBirds(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
