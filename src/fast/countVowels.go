/*
https://www.codewars.com/kata/54ff3102c1bad923760001f3
Vowel Count

Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/

package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	// Enter solution here

	splited := strings.Split(str, "")

	for i := 0; i < len(splited); i++ {
		if splited[i] == "a" || splited[i] == "e" || splited[i] == "i" || splited[i] == "o" || splited[i] == "u" {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
}
