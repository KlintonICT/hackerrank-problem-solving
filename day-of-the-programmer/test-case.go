package day_of_the_programmer

type TestCase struct {
	id       int
	data     int32
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     2017,
			expected: "13.09.2017",
		},
		{
			id:       2,
			data:     2016,
			expected: "12.09.2016",
		},
		{
			id:       3,
			data:     1800,
			expected: "12.09.1800",
		},
	}
}
