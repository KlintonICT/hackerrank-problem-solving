package solve_me_first

type Data struct {
	a uint32
	b uint32
}

type TestCase struct {
	id       int
	data     Data
	expected uint32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{a: 7, b: 3},
			expected: 10,
		},
		{
			id:       2,
			data:     Data{a: 2, b: 3},
			expected: 5,
		},
	}
}
