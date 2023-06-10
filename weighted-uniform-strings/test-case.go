package weighted_uniform_strings

type Data struct {
	s       string
	queries []int32
}

type TestCase struct {
	id       int
	data     Data
	expected []string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{s: "abbcccdddd", queries: []int32{1, 7, 5, 4, 15}},
			expected: []string{"Yes", "No", "No", "Yes", "No"},
		},
		{
			id:       2,
			data:     Data{s: "abccddde", queries: []int32{1, 3, 12, 5, 9, 10}},
			expected: []string{"Yes", "Yes", "Yes", "Yes", "No", "No"},
		},
		{
			id:       3,
			data:     Data{s: "aaabbbbcccddd", queries: []int32{9, 7, 8, 12, 5}},
			expected: []string{"Yes", "No", "Yes", "Yes", "No"},
		},
	}
}
