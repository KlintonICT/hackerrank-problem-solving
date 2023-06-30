package day_of_the_programmer

import (
	"fmt"
	"strconv"
)

func dayOfProgrammer(year int32) string {
	isLeapYear := false

	if year < 1918 {
		isLeapYear = year%4 == 0
	} else if year > 1918 {
		isLeapYear = (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	} else {
		return "26.09.1918"
	}

	if isLeapYear {
		return "12.09." + strconv.Itoa(int(year))
	}

	return "13.09." + strconv.Itoa(int(year))
}

func Run() {
	tc := &TestCase{}

	fmt.Println("Day of the programmer")

	for _, tc := range tc.getTestCases() {
		result := dayOfProgrammer(tc.data)

		fmt.Printf("TestCase: %d; data: %d; result: %s\n", tc.id, tc.data, result)
	}
}
