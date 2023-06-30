package breaking_the_records

type TestCase struct {
	id       int
	data     []int32
	expected []int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{10, 5, 20, 20, 4, 5, 2, 25, 1},
			expected: []int32{2, 4},
		},
		{
			id:       2,
			data:     []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42},
			expected: []int32{4, 0},
		},
	}
}
