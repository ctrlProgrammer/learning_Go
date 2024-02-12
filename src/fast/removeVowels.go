/**
https://www.codewars.com/kata/52fba66badcd10859f00097e/go
Disemvowel Trolls

Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.
*/

package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	newString := ""
	spl := strings.Split(comment, "")

	for i := 0; i < len(spl); i++ {
		isVowel := spl[i] == "a" || spl[i] == "e" || spl[i] == "i" || spl[i] == "o" || spl[i] == "u" || spl[i] == "A" || spl[i] == "E" || spl[i] == "I" || spl[i] == "O" || spl[i] == "U"

		if !isVowel {
			newString += spl[i]
		}
	}

	return newString
}

func main() {
	fmt.Println(Disemvowel("Hello comments"))
}
