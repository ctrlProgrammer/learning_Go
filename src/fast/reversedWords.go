/*
CodeWards https://www.codewars.com/kata/5264d2b162488dc400000001
Stop gninnipS My sdroW!

Write a function that takes in a string of one or more words, and returns the same string, but with all words that have five or more letters reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

Examples:

"Hey fellow warriors"  --> "Hey wollef sroirraw"
"This is a test        --> "This is a test"
"This is another test" --> "This is rehtona test"
*/

package kata

import (
	"strings"
)

func SpinWords(str string) string {
	newStr := ""
	words := strings.Split(str, " ")

	if len(words) > 0 {
		for i := 0; i < len(words); i++ {
			letters := strings.Split(words[i], "")
			var parsedAddition string

			if len(letters) >= 5 {
				newWord := letters

				for i, j := 0, len(newWord)-1; i < j; i, j = i+1, j-1 {
					newWord[i], newWord[j] = newWord[j], newWord[i]
				}

				parsedAddition = strings.Join(newWord, "")
			} else {
				parsedAddition = words[i]
			}

			if i < len(words)-1 {
				newStr += parsedAddition + " "
			} else {
				newStr += parsedAddition
			}
		}
	}

	return newStr
} // SpinWords
