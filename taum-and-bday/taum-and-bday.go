package taum_and_bday

import "fmt"

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	blackCost := wc + z
	whiteCost := bc + z

	if bc < blackCost {
		blackCost = bc
	}

	if wc < whiteCost {
		whiteCost = wc
	}

	return int64(b)*int64(blackCost) + int64(w)*int64(whiteCost)
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Taum and B'day")

	for _, tc := range tc.getTestCases() {
		result := taumBday(tc.data.b, tc.data.w, tc.data.bc, tc.data.wc, tc.data.z)

		fmt.Printf("TestCase: %d; data: %+v; result: %d\n", tc.id, tc.data, result)
	}
}
