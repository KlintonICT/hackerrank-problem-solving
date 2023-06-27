package birthday_cake_candles

type TestCase struct {
	id       int
	data     []int32
	expected int32
}

func (tc *TestCase) getTestCases() []TestCase {
	return []TestCase{
		{id: 1, data: []int32{4, 4, 1, 3}, expected: 2},
		{id: 2, data: []int32{3, 2, 1, 3}, expected: 2},
		{id: 3, data: []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}, expected: 5},
	}
}
