package picking_numbers

type TestCase struct {
	id       int
	data     []int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{1, 1, 2, 2, 4, 4, 5, 5, 5},
			expected: 5,
		},
		{
			id:       2,
			data:     []int32{4, 6, 5, 3, 3, 1},
			expected: 3,
		},
		{
			id:       3,
			data:     []int32{1, 2, 2, 3, 1, 2},
			expected: 5,
		},
	}
}
