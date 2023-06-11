package main

import (
	"fmt"
	"os"
	angry_professor "problem-solving/angry-professor"
	circular_array_rotation "problem-solving/circular-array-rotation"
	minimum_absolute_difference "problem-solving/minimum-absolute-difference"
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
