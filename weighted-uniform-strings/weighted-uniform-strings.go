package weighted_uniform_strings

import "fmt"

func weightedUniformStrings(s string, queries []int32) []string {
	weights := make(map[rune]int)
	uniformWeights := make(map[string]int)

	var result []string

	for char, weight := 'a', 1; char <= 'z'; char, weight = char+1, weight+1 {
		weights[char] = weight
	}

	chars := string(s[0])
	charsWeight := weights[rune(s[0])]
	uniformWeights[chars] = charsWeight
	hasValue := false

	for i := 0; i < len(s); i++ {
		if i > 0 {
			currentChar := s[i]
			previousChar := s[i-1]

			if currentChar == previousChar {
				charsWeight += weights[rune(currentChar)]
				chars += string(currentChar)
			} else {
				chars = string(currentChar)
				charsWeight = weights[rune(currentChar)]
			}
		}

		for index, q := range queries {
			if !hasValue {
				if q == int32(charsWeight) {
					result = append(result, "Yes")
				} else {
					result = append(result, "No")
				}
			} else {
				if q == int32(charsWeight) {
					result[index] = "Yes"
				}
			}
		}

		hasValue = true
	}

	return result
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Weight Uniform Strings")

	for _, tc := range tc.getTestCases() {
		result := weightedUniformStrings(tc.data.s, tc.data.queries)

		fmt.Printf("TestCase: %d; data: %+v; result: %v\n", tc.id, tc.data, result)
	}
}
