package main

import (
	"fmt"
	circular_array_rotation "problem-solving/circular-array-rotation"
	"strings"
)

func main() {
	divider := strings.Repeat("=", 100)
	solutionFuncs := []func(){circular_array_rotation.Run}

	for _, solutionFunc := range solutionFuncs {
		fmt.Println(divider)
		solutionFunc()
		fmt.Println(divider)
	}
}
