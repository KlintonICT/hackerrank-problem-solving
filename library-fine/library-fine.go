package library_fine

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	diffY := y1 - y2
	if diffY > 0 {
		return 10000
	}

	diffM := m1 - m2
	if diffY == 0 && diffM > 0 {
		return 500 * diffM
	}

	diffD := d1 - d2
	if diffY == 0 && diffM == 0 && diffD > 0 {
		return 15 * diffD
	}

	return 0
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Library Fine")

	for _, tc := range tc.getTestCases() {
		result := libraryFine(tc.data.d1, tc.data.m1, tc.data.y1, tc.data.d2, tc.data.m2, tc.data.y2)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
