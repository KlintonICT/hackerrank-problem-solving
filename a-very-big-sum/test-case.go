package a_very_big_sum

type TestCase struct {
	id       int
	data     []int64
	expected int64
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005},
			expected: 5000000015,
		},
	}
}
