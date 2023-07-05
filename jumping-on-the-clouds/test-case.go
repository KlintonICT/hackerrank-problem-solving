package jumping_on_the_clouds

type TestCase struct {
	id       int
	data     []int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{0, 1, 0, 0, 0, 1, 0},
			expected: 3,
		},
		{
			id:       2,
			data:     []int32{0, 0, 1, 0, 0, 1, 0},
			expected: 4,
		},
		{
			id:       3,
			data:     []int32{0, 0, 0, 0, 1, 0},
			expected: 3,
		},
	}
}
