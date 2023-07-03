package jumping_on_the_clouds_revisited

type Data struct {
	c []int32
	k int32
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
			data:     Data{c: []int32{0, 0, 1, 0}, k: 2},
			expected: 96,
		},
		{
			id:       2,
			data:     Data{c: []int32{0, 0, 1, 0, 0, 1, 1, 0}, k: 2},
			expected: 92,
		},
	}
}
