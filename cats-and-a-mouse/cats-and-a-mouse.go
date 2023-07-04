package cats_and_a_mouse

import (
	"fmt"
	"math"
)

func catAndMouse(x int32, y int32, z int32) string {
	xToZ := math.Abs(float64(x - z))
	ytoZ := math.Abs(float64(y - z))

	if xToZ < ytoZ {
		return "Cat A"
	} else if xToZ > ytoZ {
		return "Cat B"
	}

	return "Mouse C"
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Cats and a Mouse")

	for _, tc := range tc.getTestCases() {
		result := catAndMouse(tc.data.x, tc.data.y, tc.data.z)

		fmt.Printf("TestCase: %d; data: %+v; result: %s\n", tc.id, tc.data, result)
	}
}
