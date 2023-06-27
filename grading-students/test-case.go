package grading_students

type TestCase struct {
	id       int
	data     []int32
	expected []int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{84, 29, 57},
			expected: []int32{85, 29, 57},
		},
		{
			id:       2,
			data:     []int32{73, 67, 38, 33},
			expected: []int32{75, 67, 40, 33},
		},
	}
}
