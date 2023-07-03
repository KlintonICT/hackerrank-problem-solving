package find_digits

type TestCase struct {
	id       int
	data     int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     124,
			expected: 3,
		},
		{
			id:       2,
			data:     111,
			expected: 3,
		},
		{
			id:       3,
			data:     10,
			expected: 1,
		},
		{
			id:       4,
			data:     12,
			expected: 2,
		},
		{
			id:       5,
			data:     1012,
			expected: 3,
		},
	}
}
