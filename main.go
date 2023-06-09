package main

import (
	"fmt"
	angry_professor "problem-solving/angry-professor"
	circular_array_rotation "problem-solving/circular-array-rotation"
	"strings"
)

func main() {
	divider := strings.Repeat("=", 100)
	solutionFuncs := []func(){circular_array_rotation.Run, angry_professor.Run}

	for _, solutionFunc := range solutionFuncs {
		fmt.Println(divider)
		solutionFunc()
	}
	fmt.Println(divider)
}
