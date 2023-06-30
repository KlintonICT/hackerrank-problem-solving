package drawing_book

type Data struct {
	n int32
	p int32
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
			data:     Data{n: 5, p: 3},
			expected: 1,
		},
		{
			id:       2,
			data:     Data{n: 6, p: 2},
			expected: 1,
		},
		{
			id:       3,
			data:     Data{n: 5, p: 4},
			expected: 0,
		},
	}
}
