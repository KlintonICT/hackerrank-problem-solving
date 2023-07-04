package library_fine

type Data struct {
	d1 int32
	m1 int32
	y1 int32
	d2 int32
	m2 int32
	y2 int32
}

type TestCase struct {
	id       int
	data     Data
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{d1: 14, m1: 7, y1: 2018, d2: 5, m2: 7, y2: 2018},
			expected: 135,
		},
		{
			id:       2,
			data:     Data{d1: 9, m1: 6, y1: 2015, d2: 6, m2: 6, y2: 2015},
			expected: 45,
		},
		{
			id:       3,
			data:     Data{d1: 2, m1: 7, y1: 1014, d2: 1, m2: 1, y2: 1015},
			expected: 0,
		},
	}
}
