package hackerrank_in_a_string

type TestCase struct {
	id       int
	data     string
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     "hereiamstackerrank",
			expected: "YES",
		},
		{
			id:       2,
			data:     "hackerworld",
			expected: "NO",
		},
		{
			id:       3,
			data:     "hhaacckkekraraannk",
			expected: "YES",
		},
		{
			id:       4,
			data:     "rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt",
			expected: "NO",
		},
	}
}
