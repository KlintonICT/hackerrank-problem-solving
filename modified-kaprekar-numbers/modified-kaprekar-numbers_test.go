package modified_kaprekar_numbers

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestKaprekarNumbers(t *testing.T) {
	tc := &TestCase{}

	for _, tc := range tc.getTestCases() {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		kaprekarNumbers(tc.data.p, tc.data.q)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout

		if string(out) != tc.expected {
			t.Errorf("TestCase: %d; actual: %s; expected: %s", tc.id, string(out), tc.expected)
		}
	}
}
