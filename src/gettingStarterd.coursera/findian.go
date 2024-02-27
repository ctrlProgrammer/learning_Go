package main

import (
	"fmt"
	"strings"
)

func main() {
	var userString string
	var validatedLetters [3]int

	validatedLetters[0] = 0
	validatedLetters[1] = 0
	validatedLetters[2] = 0

	fmt.Printf("Write a word: ")
	fmt.Scan(&userString)

	splitedString := strings.Split(userString, "")

	for i := 0; i < len(splitedString); i++ {

		switch splitedString[i] {
		case "i", "I":
			validatedLetters[0]++
		case "a", "A":
			validatedLetters[1]++
		case "n", "N":
			validatedLetters[2]++
		default:
		}
	}

	fmt.Println("Found Is on the word", validatedLetters[0])
	fmt.Println("Found As on the word", validatedLetters[1])
	fmt.Println("Found Ns on the word", validatedLetters[2])

	if validatedLetters[0] > 0 && validatedLetters[1] > 0 && validatedLetters[2] > 0 {
		fmt.Println("Found IAN on your word. Found!")
	} else {
		if validatedLetters[0] > 0 {
			fmt.Println("Only found I on your word. Not found!")
		} else if validatedLetters[1] > 0 {
			fmt.Println("Only found A on your word. Not found!")
		} else if validatedLetters[2] > 0 {
			fmt.Println("Only found N on your word. Not found!")
		} else {
			fmt.Println("Not found any IAN letters on the word. Not found!")
		}
	}
}
