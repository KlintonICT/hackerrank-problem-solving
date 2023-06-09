package circular_array_rotation

type Data struct {
	a       []int32
	k       int32
	queries []int32
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
			data:     Data{a: []int32{3, 4, 5}, k: int32(2), queries: []int32{1, 2}},
			expected: []int32{5, 3},
		},
		{
			id:       2,
			data:     Data{a: []int32{1, 2, 3}, k: int32(2), queries: []int32{0, 1, 2}},
			expected: []int32{2, 3, 1},
		},
	}
}
