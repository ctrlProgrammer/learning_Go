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
