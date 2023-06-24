package plus_minus

type TestCase struct {
	id       int
	data     []int32
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{1, 1, 0, -1, -1},
			expected: "0.400000\n0.400000\n0.200000\n",
		},
		{
			id:       2,
			data:     []int32{-4, 3, -9, 0, 4, 1},
			expected: "0.500000\n0.333333\n0.166667\n",
		},
	}
}
