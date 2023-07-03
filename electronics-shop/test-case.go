package electronics_shop

type Data struct {
	keyboards []int32
	drives    []int32
	b         int32
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
			data:     Data{keyboards: []int32{40, 50, 60}, drives: []int32{5, 8, 12}, b: 60},
			expected: 58,
		},
		{
			id:       2,
			data:     Data{keyboards: []int32{3, 1}, drives: []int32{5, 2, 8}, b: 10},
			expected: 9,
		},
		{
			id:       3,
			data:     Data{keyboards: []int32{4}, drives: []int32{5}, b: 5},
			expected: -1,
		},
	}
}
