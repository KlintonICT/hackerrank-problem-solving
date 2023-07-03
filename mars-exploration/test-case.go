package mars_exploration

type TestCase struct {
	id       int
	data     string
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     "SOSTOT",
			expected: 2,
		},
		{
			id:       2,
			data:     "SOSSPSSQSSOR",
			expected: 3,
		},
		{
			id:       3,
			data:     "SOSSOT",
			expected: 1,
		},
		{
			id:       4,
			data:     "SOSSOSSOS",
			expected: 0,
		},
	}
}
