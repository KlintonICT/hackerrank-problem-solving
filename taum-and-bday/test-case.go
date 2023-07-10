package taum_and_bday

type Data struct {
	b  int32
	w  int32
	bc int32
	wc int32
	z  int32
}

type TestCase struct {
	id       int
	data     Data
	expected int64
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{b: 3, w: 5, bc: 3, wc: 4, z: 1},
			expected: 29,
		},
		{
			id:       2,
			data:     Data{b: 10, w: 10, bc: 1, wc: 1, z: 1},
			expected: 20,
		},
		{
			id:       3,
			data:     Data{b: 5, w: 9, bc: 2, wc: 3, z: 4},
			expected: 37,
		},
		{
			id:       4,
			data:     Data{b: 3, w: 6, bc: 9, wc: 1, z: 1},
			expected: 12,
		},
		{
			id:       5,
			data:     Data{b: 7, w: 7, bc: 4, wc: 2, z: 1},
			expected: 35,
		},
		{
			id:       6,
			data:     Data{b: 3, w: 3, bc: 1, wc: 9, z: 2},
			expected: 12,
		},
		{
			id:       7,
			data:     Data{b: 27984, w: 1402, bc: 619246, wc: 615589, z: 247954},
			expected: 18192035842,
		},
	}
}
