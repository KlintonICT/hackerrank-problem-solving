package divisible_sum_pairs

type Data struct {
	n  int32
	k  int32
	ar []int32
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
			data:     Data{n: 6, k: 3, ar: []int32{1, 3, 2, 6, 1, 2}},
			expected: 5,
		},
	}
}
