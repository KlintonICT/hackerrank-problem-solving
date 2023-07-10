package modified_kaprekar_numbers

type Data struct {
	p int32
	q int32
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
			data:     Data{p: 1, q: 100},
			expected: "1 9 45 55 99\n",
		},
	}
}
