package beautiful_days_at_the_movies

type Data struct {
	i int32
	j int32
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
			data:     Data{i: 20, j: 23, k: 6},
			expected: 2,
		},
	}
}
