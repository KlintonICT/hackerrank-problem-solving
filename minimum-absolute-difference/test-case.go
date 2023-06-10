package minimum_absolute_difference

type TestCase struct {
	id       int
	data     []int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{-2, 2, 4},
			expected: 2,
		},
		{
			id:       2,
			data:     []int32{-3, 7, 0},
			expected: 3,
		},
		{
			id:       3,
			data:     []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75},
			expected: 1,
		},
		{
			id:       4,
			data:     []int32{1, -3, 71, 68, 17},
			expected: 3,
		},
	}
}
