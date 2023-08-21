package beautiful_triplets

type Data struct {
	d   int32
	arr []int32
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
			data:     Data{d: 1, arr: []int32{2, 2, 3, 4, 5}},
			expected: 3,
		},
		{
			id:       2,
			data:     Data{d: 3, arr: []int32{1, 2, 4, 5, 7, 8, 10}},
			expected: 3,
		},
	}
}
