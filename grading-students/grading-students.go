package grading_students

import "fmt"

func gradingStudents(grades []int32) []int32 {
	finalGrades := make([]int32, len(grades))
	copy(finalGrades, grades)

	for i := 0; i < len(grades); i++ {
		if grades[i] >= 38 {
			remainder := grades[i] % 5
			if remainder >= 3 {
				finalGrades[i] = grades[i] + (5 - remainder)
			}
		}
	}

	return finalGrades
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Grading Students")

	for _, tc := range tc.getTestCases() {
		result := gradingStudents(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %v\n", tc.id, tc.data, result)
	}
}
