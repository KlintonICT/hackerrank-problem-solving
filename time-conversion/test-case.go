package time_conversion

type TestCase struct {
	id       int
	data     string
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     "12:01:00PM",
			expected: "12:01:00",
		},
		{
			id:       2,
			data:     "12:01:00AM",
			expected: "00:01:00",
		},
		{
			id:       3,
			data:     "07:05:45PM",
			expected: "19:05:45",
		},
	}
}
