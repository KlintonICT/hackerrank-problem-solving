package apple_and_orange

type Data struct {
	s       int32
	t       int32
	a       int32
	b       int32
	apples  []int32
	oranges []int32
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
			data:     Data{s: 7, t: 10, a: 4, b: 12, apples: []int32{2, 3, -4}, oranges: []int32{3, -2, -4}},
			expected: "1\n2\n",
		},
		{
			id:       2,
			data:     Data{s: 7, t: 11, a: 5, b: 15, apples: []int32{-2, 2, 1}, oranges: []int32{5, -6}},
			expected: "1\n1\n",
		},
	}
}
