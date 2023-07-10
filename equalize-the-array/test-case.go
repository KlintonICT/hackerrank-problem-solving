package equalize_the_array

type TestCase struct {
	id       int
	data     []int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{1, 2, 2, 3},
			expected: 2,
		},
		{
			id:       2,
			data:     []int32{3, 3, 2, 1, 3},
			expected: 2,
		},
	}
}
