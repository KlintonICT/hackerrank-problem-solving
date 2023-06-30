package sales_by_match

type Data struct {
	n  int32
	ar []int32
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
			data:     Data{n: 9, ar: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}},
			expected: 3,
		},
	}
}
