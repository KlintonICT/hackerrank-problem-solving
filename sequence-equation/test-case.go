package sequence_equation

type TestCase struct {
	id       int
	data     []int32
	expected []int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{5, 2, 1, 3, 4},
			expected: []int32{4, 2, 5, 1, 3},
		},
		{
			id:       2,
			data:     []int32{2, 3, 1},
			expected: []int32{2, 3, 1},
		},
		{
			id:       3,
			data:     []int32{4, 3, 5, 1, 2},
			expected: []int32{1, 3, 5, 4, 2},
		},
	}
}
