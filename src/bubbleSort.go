package main

import (
	"fmt"
	"strconv"
)

func Swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
}

func BubbleSort(numbersSlice []int) {
	var swapped bool
	var n = len(numbersSlice)

	for i := 0; i < n-1; i++ {
		swapped = false

		for j := 0; j < n-i-1; j++ {
			if numbersSlice[j] > numbersSlice[j+1] {
				Swap(&numbersSlice[j], &numbersSlice[j+1])
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}

}

func main() {
	newSlice := []int{}

	fmt.Println("Welcome")
	fmt.Println("Press X to exit")
	fmt.Println("Press A to sort your integers")
	fmt.Println("Enter a serie of integers: ")

	for {
		var userInput string

		fmt.Println("Enter a new integer (Enter A to sort your integers): ")
		fmt.Scan(&userInput)

		if userInput == "x" || userInput == "X" {
			break
		}

		if userInput == "a" || userInput == "A" {
			fmt.Println("Your sorted integers are: ")
			BubbleSort(newSlice)
			fmt.Println(newSlice)
			break
		}

		parsedValue, err := strconv.Atoi(userInput)

		if err == nil {
			newSlice = append(newSlice, parsedValue)
			fmt.Println("Your integers are: ")
			fmt.Println(newSlice)
		} else {
			fmt.Println("Invalid Number, enter an integer")
		}

	}
}
