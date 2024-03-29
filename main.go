package main

import (
	"fmt"
	"os"
	a_very_big_sum "problem-solving/a-very-big-sum"
	acm_icpc_team "problem-solving/acm-icpc-team"
	angry_professor "problem-solving/angry-professor"
	append_and_delete "problem-solving/append-and-delete"
	apple_and_orange "problem-solving/apple-and-orange"
	beautiful_days_at_the_movies "problem-solving/beautiful-days-at-the-movies"
	beautiful_triplets "problem-solving/beautiful-triplets"
	between_two_sets "problem-solving/between-two-sets"
	big_sorting "problem-solving/big-sorting"
	bill_division "problem-solving/bill-division"
	birthday_cake_candles "problem-solving/birthday-cake-candles"
	breaking_the_records "problem-solving/breaking-the-records"
	cats_and_a_mouse "problem-solving/cats-and-a-mouse"
	circular_array_rotation "problem-solving/circular-array-rotation"
	compare_the_triplets "problem-solving/compare-the-triplets"
	counting_valleys "problem-solving/counting-valleys"
	cut_the_sticks "problem-solving/cut-the-sticks"
	day_of_the_programmer "problem-solving/day-of-the-programmer"
	designer_pdf_viewer "problem-solving/designer-pdf-viewer"
	diagonal_difference "problem-solving/diagonal-difference"
	divisible_sum_pairs "problem-solving/divisible-sum-pairs"
	drawing_book "problem-solving/drawing-book"
	electronics_shop "problem-solving/electronics-shop"
	equalize_the_array "problem-solving/equalize-the-array"
	find_digits "problem-solving/find-digits"
	forming_a_magic_square "problem-solving/forming-a-magic-square"
	grading_students "problem-solving/grading-students"
	hackerrank_in_a_string "problem-solving/hackerrank-in-a-string"
	into_to_tutorial_challenges "problem-solving/intro-to-tutorial-challenges"
	jumping_on_the_clouds "problem-solving/jumping-on-the-clouds"
	jumping_on_the_clouds_revisited "problem-solving/jumping-on-the-clouds-revisited"
	library_fine "problem-solving/library-fine"
	mars_exploration "problem-solving/mars-exploration"
	migratory_birds "problem-solving/migratory-birds"
	mini_max_sum "problem-solving/mini-max-sum"
	minimum_absolute_difference "problem-solving/minimum-absolute-difference"
	modified_kaprekar_numbers "problem-solving/modified-kaprekar-numbers"
	number_line_jumps "problem-solving/number-line-jumps"
	pangrams "problem-solving/pangrams"
	picking_numbers "problem-solving/picking-numbers"
	plus_minus "problem-solving/plus-minus"
	repeated_string "problem-solving/repeated-string"
	sales_by_match "problem-solving/sales-by-match"
	save_the_prisoner "problem-solving/save-the-prisoner"
	sequence_equation "problem-solving/sequence-equation"
	sherlock_and_squares "problem-solving/sherlock-and-squares"
	simple_array_sum "problem-solving/simple-array-sum"
	solve_me_first "problem-solving/solve-me-first"
	staircase "problem-solving/staircase"
	subarray_division "problem-solving/subarray-division"
	taum_and_bday "problem-solving/taum-and-bday"
	the_hurdle_race "problem-solving/the-hurdle-race"
	time_conversion "problem-solving/time-conversion"
	utopian_tree "problem-solving/utopian-tree"
	viral_advertising "problem-solving/viral-advertising"
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
		32: the_hurdle_race.Run,
		33: designer_pdf_viewer.Run,
		34: utopian_tree.Run,
		35: beautiful_days_at_the_movies.Run,
		36: viral_advertising.Run,
		37: save_the_prisoner.Run,
		38: jumping_on_the_clouds_revisited.Run,
		39: big_sorting.Run,
		40: into_to_tutorial_challenges.Run,
		41: mars_exploration.Run,
		42: hackerrank_in_a_string.Run,
		43: pangrams.Run,
		44: sequence_equation.Run,
		45: find_digits.Run,
		46: append_and_delete.Run,
		47: sherlock_and_squares.Run,
		48: library_fine.Run,
		49: cut_the_sticks.Run,
		50: repeated_string.Run,
		51: jumping_on_the_clouds.Run,
		52: equalize_the_array.Run,
		53: acm_icpc_team.Run,
		54: taum_and_bday.Run,
		55: modified_kaprekar_numbers.Run,
		56: beautiful_triplets.Run,
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
