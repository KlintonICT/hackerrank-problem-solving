package save_the_prisoner

type Data struct {
	n int32
	m int32
	s int32
}

type TestCase struct {
	id       int
	data     Data
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{n: 4, m: 6, s: 2},
			expected: 3,
		},
		{
			id:       2,
			data:     Data{n: 5, m: 2, s: 1},
			expected: 2,
		},
		{
			id:       3,
			data:     Data{n: 5, m: 2, s: 2},
			expected: 3,
		},
		{
			id:       4,
			data:     Data{n: 7, m: 19, s: 2},
			expected: 6,
		},
		{
			id:       5,
			data:     Data{n: 3, m: 7, s: 3},
			expected: 3,
		},
	}
}
