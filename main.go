package main

import (
	"fmt"
	"os"
	a_very_big_sum "problem-solving/a-very-big-sum"
	angry_professor "problem-solving/angry-professor"
	circular_array_rotation "problem-solving/circular-array-rotation"
	compare_the_triplets "problem-solving/compare-the-triplets"
	minimum_absolute_difference "problem-solving/minimum-absolute-difference"
	simple_array_sum "problem-solving/simple-array-sum"
	solve_me_first "problem-solving/solve-me-first"
	weighted_uniform_strings "problem-solving/weighted-uniform-strings"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	divider := strings.Repeat("=", 100)

	solutionFuncs := map[int]func(){
		1: circular_array_rotation.Run,
		2: angry_professor.Run,
		3: minimum_absolute_difference.Run,
		4: weighted_uniform_strings.Run,
		5: solve_me_first.Run,
		6: simple_array_sum.Run,
		7: compare_the_triplets.Run,
		8: a_very_big_sum.Run,
	}

	if len(args) > 0 {
		problemId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("\"problem-id\" needs to be positive integer number")
		} else {
			solutionFunc, isExist := solutionFuncs[problemId]
			if isExist {
				solutionFunc()
			}
		}
	} else {
		for _, solutionFunc := range solutionFuncs {
			fmt.Println(divider)
			solutionFunc()
		}
		fmt.Println(divider)
	}
}
