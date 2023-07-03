package into_to_tutorial_challenges

type Data struct {
	V   int32
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
			data:     Data{V: 3, arr: []int32{1, 2, 3}},
			expected: 2,
		},
		{
			id:       1,
			data:     Data{V: 4, arr: []int32{1, 4, 5, 7, 9, 12}},
			expected: 1,
		},
	}
}
