package apple_and_orange

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCountApplesAndOranges(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		countApplesAndOranges(tc.data.s, tc.data.t, tc.data.a, tc.data.b, tc.data.apples, tc.data.oranges)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		if string(out) != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, string(out), tc.expected)
		}
	}
}
