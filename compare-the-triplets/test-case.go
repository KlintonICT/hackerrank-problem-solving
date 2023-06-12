package compare_the_triplets

type Data struct {
	a []int32
	b []int32
}

type TestCase struct {
	id       int
	data     Data
	expected []int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{a: []int32{1, 2, 3}, b: []int32{3, 2, 1}},
			expected: []int32{1, 1},
		},
		{
			id:       2,
			data:     Data{a: []int32{5, 6, 7}, b: []int32{3, 6, 10}},
			expected: []int32{1, 1},
		},
		{
			id:       3,
			data:     Data{a: []int32{17, 28, 30}, b: []int32{99, 16, 8}},
			expected: []int32{2, 1},
		},
	}
}
