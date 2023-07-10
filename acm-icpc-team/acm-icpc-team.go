package acm_icpc_team

import "fmt"

func acmTeam(topic []string) []int32 {
	max, countMax := int32(0), int32(0)

	for i := 0; i < len(topic)-1; i++ {
		for j := i + 1; j < len(topic); j++ {
			t1 := topic[i]
			t2 := topic[j]
			count := int32(0)

			for k := 0; k < len(t1); k++ {
				if t1[k] == '1' || t2[k] == '1' {
					count += 1
				}
			}

			if count == max {
				countMax += 1
			} else if count > max {
				countMax = 1
				max = count
			}
		}
	}

	return []int32{max, countMax}
}

func Run() {
	tc := &TestCase{}

	fmt.Println("ACM ICPC Team")

	for _, tc := range tc.getTestCases() {
		result := acmTeam(tc.data)

		fmt.Printf("TestCase: %d; data: %v; result: %d\n", tc.id, tc.data, result)
	}
}
