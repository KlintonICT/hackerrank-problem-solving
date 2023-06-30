package counting_valleys

type Data struct {
	steps int32
	path  string
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
			data:     Data{steps: 8, path: "UDDDUDUU"},
			expected: 1,
		},
	}
}
