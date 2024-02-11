package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userString string
	var startsWithI bool
	var endsWithN bool
	var containsA bool

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write a word: ")
	scanner.Scan()
	userString = scanner.Text()
	fmt.Println("captured:", userString)

	splitedString := strings.Split(userString, "")

	fmt.Println(splitedString)

	if splitedString[0] == "i" || splitedString[0] == "I" {
		startsWithI = true
	}

	if splitedString[len(splitedString)-1] == "n" || splitedString[len(splitedString)-1] == "N" {
		endsWithN = true
	}

	for i := 0; i < len(splitedString); i++ {
		if splitedString[i] == "a" || splitedString[i] == "A" {
			containsA = true
			break
		}
	}

	if startsWithI {
		fmt.Println("Start with I")
	} else {
		fmt.Println("Does not start with I")
	}

	if endsWithN {
		fmt.Println("Ends with N")
	} else {
		fmt.Println("Does not ends with N")
	}

	if containsA {
		fmt.Println("Contains an A")
	} else {
		fmt.Println("Does not contain any A")
	}

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
