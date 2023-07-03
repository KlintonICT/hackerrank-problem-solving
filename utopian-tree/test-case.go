package utopian_tree

type TestCase struct {
	id       int
	data     int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     5,
			expected: 14,
		},
		{
			id:       2,
			data:     0,
			expected: 1,
		},
		{
			id:       3,
			data:     1,
			expected: 2,
		},
		{
			id:       4,
			data:     4,
			expected: 7,
		},
	}
}
