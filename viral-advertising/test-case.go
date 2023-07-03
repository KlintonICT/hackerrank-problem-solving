package viral_advertising

type TestCase struct {
	id       int
	data     int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     5,
			expected: 24,
		},
		{
			id:       2,
			data:     3,
			expected: 9,
		},
	}
}
