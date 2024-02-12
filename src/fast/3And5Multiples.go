/*
https://www.codewars.com/kata/514b92a657cdc65150000006
Multiples of 3 or 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.

Note: If the number is a multiple of both 3 and 5, only count it once.
*/

package main

import (
	"fmt"
)

func IsMultiple3And5(number int) bool {
	return number%3 == 0 || number%5 == 0
}

func Multiple3And5(number int) int {
	sum := 0

	for i := 1; i < number; i++ {
		if IsMultiple3And5(i) {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(Multiple3And5(10))
}
