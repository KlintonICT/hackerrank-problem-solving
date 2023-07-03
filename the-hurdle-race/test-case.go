package the_hurdle_race

type Data struct {
	k      int32
	height []int32
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
			data:     Data{k: 1, height: []int32{1, 2, 3, 3, 2}},
			expected: 2,
		},
		{
			id:       2,
			data:     Data{k: 4, height: []int32{1, 6, 3, 5, 2}},
			expected: 2,
		},
		{
			id:       3,
			data:     Data{k: 7, height: []int32{2, 5, 4, 5, 2}},
			expected: 0,
		},
	}
}
