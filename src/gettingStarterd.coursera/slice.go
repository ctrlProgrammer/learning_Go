package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbersSlice := []int{}

	for {
		var userInput string

		fmt.Printf("Enter a new Integer or press X to exit: ")
		fmt.Scan(&userInput)

		if userInput == "x" || userInput == "X" {
			break
		} else {
			s, err := strconv.Atoi(userInput)

			if err != nil {
				fmt.Println("Invalid Number")
			} else {
				numbersSlice = append(numbersSlice, s)
			}
		}

		sort.Slice(numbersSlice, func(i, j int) bool {
			return numbersSlice[i] < numbersSlice[j]
		})

		fmt.Println("Your inputs are: ")
		fmt.Println(numbersSlice)
	}

}
