package number_line_jumps

import (
	"fmt"
	"math"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if (x1 > x2 && v1 >= v2) || (x2 > x1 && v2 >= v1) {
		return "NO"
	}

	timeToIntersection := int32(math.Abs(float64((x1 - x2))) / math.Abs(float64((v2 - v1))))
	meetingPoint1 := x1 + v1*timeToIntersection
	meetingPoint2 := x2 + v2*timeToIntersection

	if timeToIntersection >= 0 && meetingPoint1 == meetingPoint2 {
		return "YES"
	}

	return "NO"
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Number Line Jumps")

	for _, tc := range tc.getTestCases() {
		result := kangaroo(tc.data.x1, tc.data.v1, tc.data.x2, tc.data.v2)

		fmt.Printf("TestCase: %d; data: %+v; result: %s\n", tc.id, tc.data, result)
	}
}
