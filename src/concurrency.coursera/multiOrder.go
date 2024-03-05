package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"sync"
)

func divideMatriz(numbers []int) [][]int {
	fraction := float64(len(numbers) / 4)
	fractionSize := int64(math.Floor(fraction))
	return [][]int{numbers[:fractionSize], numbers[fractionSize : fractionSize*2], numbers[fractionSize*2 : fractionSize*3], numbers[fractionSize*3:]}
}

func sortMatriz(numbers []int) []int {
	sortedNumbers := numbers
	slices.Sort(sortedNumbers)
	fmt.Println("Sorting numbers array: ", sortedNumbers)
	return sortedNumbers
}

func combineSortedNumbers(numbers []int) []int {
	var wg sync.WaitGroup
	var sortedNumbers []int = []int{}

	fractionNumbers := divideMatriz(numbers)

	for _, matriz := range fractionNumbers {
		wg.Add(1)

		go func(matriz []int) {
			defer wg.Done()
			newSortedMatriz := sortMatriz(matriz)
			sortedNumbers = append(sortedNumbers, newSortedMatriz...)
		}(matriz)
	}

	wg.Wait()

	return sortedNumbers
}

func main() {
	var numbers []int = []int{}
	var sort bool = false

	fmt.Println("Write some integers to sort it into a 4 fractionated matrices: ")
	fmt.Println("Each 1/4 from the matriz will be sorted for a different go rutine but in the end the main routine will combine all matrices")
	fmt.Println("Press enter when you want to add another number.")
	fmt.Println("Press X when you want to exit.")
	fmt.Println("Press Y when you want to sort the numbers.")
	fmt.Println("")

	for {
		var userInput string

		fmt.Print("> ")

		_, err := fmt.Scan(&userInput)

		if err != nil {
			fmt.Println("Invalid Input, please write a valid integer.")
			return
		}

		if userInput == "x" || userInput == "X" {
			fmt.Println("Thanks for your time.")
			break
		}

		if userInput == "y" || userInput == "Y" {
			if len(numbers) > 4 {
				sort = true
				break
			} else {
				fmt.Println("You need more numbers.")
				break
			}
		}

		numberInput, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Invalid Input, please write a valid integer.")
			return
		}

		fmt.Println("Added new number: " + userInput)
		numbers = append(numbers, numberInput)
	}

	if sort {
		sortedNumbers := combineSortedNumbers(numbers)
		slices.Sort(sortedNumbers)
		fmt.Println(sortedNumbers)
	}

}
