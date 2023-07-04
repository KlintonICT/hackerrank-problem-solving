package cut_the_sticks

type TestCase struct {
	id       int
	data     []int32
	expected []int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{1, 2, 3},
			expected: []int32{3, 2, 1},
		},
		{
			id:       2,
			data:     []int32{5, 4, 4, 2, 2, 8},
			expected: []int32{6, 4, 2, 1},
		},
		{
			id:       3,
			data:     []int32{1, 2, 3, 4, 3, 3, 2, 1},
			expected: []int32{8, 6, 4, 1},
		},
	}
}
