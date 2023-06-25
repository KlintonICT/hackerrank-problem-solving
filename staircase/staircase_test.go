package staircase

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestStaircase(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		staircase(tc.data)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		if string(out) != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, string(out), tc.expected)
		}
	}
}
