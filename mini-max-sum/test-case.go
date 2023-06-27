package mini_max_sum

type TestCase struct {
	id       int
	data     []int32
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     []int32{1, 3, 5, 7, 9},
			expected: "16 24\n",
		},
		{
			id:       2,
			data:     []int32{1, 2, 3, 4, 5},
			expected: "10 14\n",
		},
		{
			id:       3,
			data:     []int32{256741038, 623958417, 467905213, 714532089, 938071625},
			expected: "2063136757 2744467344\n",
		},
	}
}
