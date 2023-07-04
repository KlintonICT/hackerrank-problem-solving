package append_and_delete

import "fmt"

func appendAndDelete(s string, t string, k int32) string {
	sLen, tLen := int32(len(s)), int32(len(t))
	minLen := int32(tLen)
	commonLen := int32(0)

	if sLen < minLen {
		minLen = sLen
	}

	for i := int32(0); i < minLen; i++ {
		if t[i] != s[i] {
			break
		}
		commonLen++
	}

	minOps := sLen - commonLen + tLen - commonLen

	if k < minOps {
		return "No"
	}

	if k >= sLen+tLen {
		return "Yes"
	}

	if (k-minOps)%2 == 0 {
		return "Yes"
	}

	return "No"
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Append and Delete")

	for _, tc := range tc.getTestCases() {
		result := appendAndDelete(tc.data.s, tc.data.t, tc.data.k)

		fmt.Printf("TestCase: %d; data: %+v; result: %s\n", tc.id, tc.data, result)
	}
}
