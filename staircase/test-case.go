package staircase

type TestCase struct {
	id       int
	data     int32
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     4,
			expected: "   #\n  ##\n ###\n####\n",
		},
		{
			id:       2,
			data:     6,
			expected: "     #\n    ##\n   ###\n  ####\n #####\n######\n",
		},
	}
}
