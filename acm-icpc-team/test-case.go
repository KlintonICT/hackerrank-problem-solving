package acm_icpc_team

type TestCase struct {
	id       int
	data     []string
	expected []int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []string{"10101", "11110", "00010"},
			expected: []int32{5, 1},
		},
		{
			id:       2,
			data:     []string{"10101", "11100", "11010", "00101"},
			expected: []int32{5, 2},
		},
	}
}
