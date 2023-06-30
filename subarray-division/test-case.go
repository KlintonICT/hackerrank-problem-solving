package subarray_division

type Data struct {
	s []int32
	d int32
	m int32
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
			data:     Data{s: []int32{1, 2, 1, 3, 2}, d: 3, m: 2},
			expected: 2,
		},
		{
			id:       2,
			data:     Data{s: []int32{1, 1, 1, 1, 1, 1}, d: 3, m: 2},
			expected: 0,
		},
		{
			id:       3,
			data:     Data{s: []int32{4}, d: 4, m: 1},
			expected: 1,
		},
		{
			id:       4,
			data:     Data{s: []int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}, d: 18, m: 7},
			expected: 3,
		},
	}
}
