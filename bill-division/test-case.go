package bill_division

type Data struct {
	bill []int32
	k    int32
	b    int32
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
			data:     Data{bill: []int32{3, 10, 2, 9}, k: 1, b: 12},
			expected: "5\n",
		},
		{
			id:       2,
			data:     Data{bill: []int32{3, 10, 2, 9}, k: 1, b: 7},
			expected: "Bon Appetit\n",
		},
	}
}
