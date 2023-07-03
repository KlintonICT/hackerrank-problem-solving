package cats_and_a_mouse

type Data struct {
	x int32
	y int32
	z int32
}

type TestCase struct {
	id       int
	data     Data
	expected string
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{
			id:       1,
			data:     Data{x: 2, y: 5, z: 4},
			expected: "Cat B",
		},
		{
			id:       2,
			data:     Data{x: 1, y: 2, z: 3},
			expected: "Cat B",
		},
		{
			id:       3,
			data:     Data{x: 1, y: 3, z: 2},
			expected: "Mouse C",
		},
	}
}
