package angry_professor

type Data struct {
	k int32
	a []int32
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
			data:     Data{k: 3, a: []int32{-2, -1, 0, 1, 2}},
			expected: "NO",
		},
		{
			id:       2,
			data:     Data{k: 3, a: []int32{-1, -3, 4, 2}},
			expected: "YES",
		},
		{
			id:       3,
			data:     Data{k: 2, a: []int32{0, -1, 2, 1}},
			expected: "NO",
		},
	}
}
