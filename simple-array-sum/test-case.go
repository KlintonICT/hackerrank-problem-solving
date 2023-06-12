package simple_array_sum

type TestCase struct {
	id       int
	data     []int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{1, 2, 3},
			expected: 6,
		},
		{
			id:       2,
			data:     []int32{1, 2, 3, 4, 10, 11},
			expected: 31,
		},
	}
}
