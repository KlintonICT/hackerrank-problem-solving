package between_two_sets

type Data struct {
	a []int32
	b []int32
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
			data:     Data{a: []int32{2, 6}, b: []int32{24, 36}},
			expected: 2,
		},
		{
			id:       2,
			data:     Data{a: []int32{2, 4}, b: []int32{16, 32, 96}},
			expected: 3,
		},
	}
}
