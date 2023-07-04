package sherlock_and_squares

type Data struct {
	a int32
	b int32
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
			data:     Data{a: 24, b: 49},
			expected: 3,
		},
		{
			id:       2,
			data:     Data{a: 3, b: 9},
			expected: 2,
		},
		{
			id:       3,
			data:     Data{a: 17, b: 24},
			expected: 0,
		},
	}
}
