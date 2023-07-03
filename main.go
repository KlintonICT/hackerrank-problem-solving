package main

import (
	"fmt"
	"os"
	a_very_big_sum "problem-solving/a-very-big-sum"
	angry_professor "problem-solving/angry-professor"
	apple_and_orange "problem-solving/apple-and-orange"
	between_two_sets "problem-solving/between-two-sets"
	bill_division "problem-solving/bill-division"
	birthday_cake_candles "problem-solving/birthday-cake-candles"
	breaking_the_records "problem-solving/breaking-the-records"
	cats_and_a_mouse "problem-solving/cats-and-a-mouse"
	circular_array_rotation "problem-solving/circular-array-rotation"
	compare_the_triplets "problem-solving/compare-the-triplets"
	counting_valleys "problem-solving/counting-valleys"
	day_of_the_programmer "problem-solving/day-of-the-programmer"
	diagonal_difference "problem-solving/diagonal-difference"
	divisible_sum_pairs "problem-solving/divisible-sum-pairs"
	drawing_book "problem-solving/drawing-book"
	electronics_shop "problem-solving/electronics-shop"
	forming_a_magic_square "problem-solving/forming-a-magic-square"
	grading_students "problem-solving/grading-students"
	migratory_birds "problem-solving/migratory-birds"
	mini_max_sum "problem-solving/mini-max-sum"
	minimum_absolute_difference "problem-solving/minimum-absolute-difference"
	number_line_jumps "problem-solving/number-line-jumps"
	picking_numbers "problem-solving/picking-numbers"
	plus_minus "problem-solving/plus-minus"
	sales_by_match "problem-solving/sales-by-match"
	simple_array_sum "problem-solving/simple-array-sum"
	solve_me_first "problem-solving/solve-me-first"
	staircase "problem-solving/staircase"
	subarray_division "problem-solving/subarray-division"
	time_conversion "problem-solving/time-conversion"
	weighted_uniform_strings "problem-solving/weighted-uniform-strings"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	divider := strings.Repeat("=", 100)

	solutionFuncs := map[int]func(){
		1:  circular_array_rotation.Run,
		2:  angry_professor.Run,
		3:  minimum_absolute_difference.Run,
		4:  weighted_uniform_strings.Run,
		5:  solve_me_first.Run,
		6:  simple_array_sum.Run,
		7:  compare_the_triplets.Run,
		8:  a_very_big_sum.Run,
		9:  diagonal_difference.Run,
		10: plus_minus.Run,
		11: between_two_sets.Run,
		12: staircase.Run,
		13: mini_max_sum.Run,
		14: birthday_cake_candles.Run,
		15: time_conversion.Run,
		16: grading_students.Run,
		17: apple_and_orange.Run,
		18: number_line_jumps.Run,
		19: breaking_the_records.Run,
		20: subarray_division.Run,
		21: divisible_sum_pairs.Run,
		22: migratory_birds.Run,
		23: day_of_the_programmer.Run,
		24: bill_division.Run,
		25: sales_by_match.Run,
		26: drawing_book.Run,
		27: counting_valleys.Run,
		28: electronics_shop.Run,
		29: cats_and_a_mouse.Run,
		30: forming_a_magic_square.Run,
		31: picking_numbers.Run,
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
