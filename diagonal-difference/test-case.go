package diagonal_difference

type TestCase struct {
	id       int
	data     [][]int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}},
			expected: 2,
		},
		{
			id:       2,
			data:     [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}},
			expected: 15,
		},
	}
}
