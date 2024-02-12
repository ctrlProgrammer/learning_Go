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
