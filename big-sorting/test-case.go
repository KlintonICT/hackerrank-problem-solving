package big_sorting

type TestCase struct {
	id       int
	data     []string
	expected []string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []string{"1", "200", "150", "3"},
			expected: []string{"1", "3", "150", "200"},
		},
		{
			id:       2,
			data:     []string{"31415926535897932384626433832795", "1", "3", "10", "3", "5"},
			expected: []string{"1", "3", "3", "5", "10", "31415926535897932384626433832795"},
		},
		{
			id:       3,
			data:     []string{"1", "2", "100", "12303479849857341718340192371", "3084193741082937", "3084193741082938", "111", "200"},
			expected: []string{"1", "2", "100", "111", "200", "3084193741082937", "3084193741082938", "12303479849857341718340192371"},
		},
	}
}
