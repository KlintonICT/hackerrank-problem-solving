package migratory_birds

type TestCase struct {
	id       int
	data     []int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{1, 4, 4, 4, 5, 3},
			expected: 4,
		},
		{
			id:       2,
			data:     []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4},
			expected: 3,
		},
	}
}
