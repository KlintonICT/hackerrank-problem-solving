package angry_professor

import "fmt"

func angryProfessor(k int32, a []int32) string {
	var onTimeStudents []int32

	for _, item := range a {
		if item <= 0 {
			onTimeStudents = append(onTimeStudents, item)
		}
	}

	if int32(len(onTimeStudents)) >= k {
		return "NO"
	}

	return "YES"
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Angry Professor")

	for _, tc := range tc.getTestCases() {
		result := angryProfessor(tc.data.k, tc.data.a)

		fmt.Printf("TestCase: %d; data: %+v; result: %v\n", tc.id, tc.data, result)
	}
}
