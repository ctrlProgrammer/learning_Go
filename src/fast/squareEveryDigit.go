/*
https://www.codewars.com/kata/546e2562b03326a88e000020
Square Every Digit

Welcome. In this kata, you are asked to square every digit of a number and concatenate them.

For example, if we run 9119 through the function, 811181 will come out, because 92 is 81 and 12 is 1. (81-1-1-81)

Example #2: An input of 765 will/should return 493625 because 72 is 49, 62 is 36, and 52 is 25. (49-36-25)

Note: The function accepts an integer and returns an integer.

Happy Coding!
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SquareAllDigits(number int) int {
	stringNumbers := strconv.Itoa(number)
	spl := strings.Split(stringNumbers, "")
	newString := ""

	for i := 0; i < len(spl); i++ {
		parsedNumber, err := strconv.Atoi(spl[i])
		if err == nil {
			newString += strconv.Itoa(parsedNumber * parsedNumber)
		}
	}

	validatedNumber, _ := strconv.Atoi(newString)

	return validatedNumber
}

func main() {
	fmt.Println(SquareAllDigits(11155236552))
}
