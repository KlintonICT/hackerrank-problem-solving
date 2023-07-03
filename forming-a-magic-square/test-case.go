package forming_a_magic_square

type TestCase struct {
	id       int
	data     [][]int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}},
			expected: 7,
		},
		{
			id:       2,
			data:     [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}},
			expected: 1,
		},
		{
			id:       3,
			data:     [][]int32{{4, 8, 2}, {4, 5, 7}, {6, 1, 6}},
			expected: 4,
		},
	}
}
