package number_line_jumps

type Data struct {
	x1 int32
	v1 int32
	x2 int32
	v2 int32
}

type TestCase struct {
	id       int
	data     Data
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{x1: 2, v1: 1, x2: 1, v2: 2},
			expected: "YES",
		},
		{
			id:       2,
			data:     Data{x1: 0, v1: 3, x2: 4, v2: 2},
			expected: "YES",
		},
		{
			id:       3,
			data:     Data{x1: 0, v1: 2, x2: 5, v2: 3},
			expected: "NO",
		},
	}
}
