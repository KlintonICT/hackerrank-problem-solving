package repeated_string

type Data struct {
	s string
	n int64
}

type TestCase struct {
	id       int
	data     Data
	expected int64
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{s: "abcac", n: 10},
			expected: 4,
		},
		{
			id:       2,
			data:     Data{s: "aba", n: 10},
			expected: 7,
		},
		{
			id:       3,
			data:     Data{s: "a", n: 1000000000000},
			expected: 1000000000000,
		},
	}
}
