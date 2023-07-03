package designer_pdf_viewer

type Data struct {
	h    []int32
	word string
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
			data:     Data{h: []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5}, word: "torn"},
			expected: 8,
		},
		{
			id:       2,
			data:     Data{h: []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, word: "abc"},
			expected: 9,
		},
		{
			id:       3,
			data:     Data{h: []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, word: "zaba"},
			expected: 28,
		},
	}
}
