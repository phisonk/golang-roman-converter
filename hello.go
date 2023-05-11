package main

import (
	"fmt"
)

var RomanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	greatest := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := RomanNumerals[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum
}
func main() {
	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your Roman Number: ")

	// var then variable name then variable type
	var roman string

	// Taking input from user
	fmt.Scanln(&roman)

	// Print function is used to
	// display output in the same line
	fmt.Printf("Your Decimal Numeral Value is: %v\n", romanToInt(roman))

}
