package pangrams

type TestCase struct {
	id       int
	data     string
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     "The quick brown fox jumped over the lazy dog",
			expected: "not pangram",
		},
		{
			id:       2,
			data:     "We promptly judged antique ivory buckles for the next prize",
			expected: "pangram",
		},
		{
			id:       3,
			data:     "We promptly judged antique ivory buckles for the prize",
			expected: "not pangram",
		},
	}
}
