package append_and_delete

type Data struct {
	s string
	t string
	k int32
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
			data:     Data{s: "abc", t: "def", k: 6},
			expected: "Yes",
		},
		{
			id:       2,
			data:     Data{s: "hackerhappy", t: "hackerrank", k: 9},
			expected: "Yes",
		},
		{
			id:       3,
			data:     Data{s: "aba", t: "aba", k: 7},
			expected: "Yes",
		},
		{
			id:       4,
			data:     Data{s: "ashley", t: "ash", k: 2},
			expected: "No",
		},
		{
			id:       5,
			data:     Data{s: "aaaaaaaaaa", t: "aaaaa", k: 7},
			expected: "Yes",
		},
		{
			id:       6,
			data:     Data{s: "zzzzz", t: "zzzzzzz", k: 4},
			expected: "Yes",
		},
		{
			id:       7,
			data:     Data{s: "y", t: "yu", k: 2},
			expected: "No",
		},
		{
			id:       8,
			data:     Data{s: "qwerasdf", t: "qwerbsdf", k: 6},
			expected: "No",
		},
		{
			id:       9,
			data:     Data{s: "asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv", t: "asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv", k: 20},
			expected: "Yes",
		},
	}
}
