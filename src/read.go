package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserData struct {
	fname string
	lname string
}

func main() {
	namesList := []UserData{}

	dat, err := os.Open("src/testText.txt")

	if err != nil {
		fmt.Println("Error reading the file. ", err)
		return
	}

	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		parsedLine := strings.Split(fileScanner.Text(), " ")
		namesList = append(namesList, UserData{fname: parsedLine[0], lname: parsedLine[1]})
	}

	dat.Close()

	for i := 0; i < len(namesList); i++ {
		fmt.Println("User Number", i+1)
		fmt.Println("Name: " + namesList[i].fname + ", Lastname: " + namesList[i].lname)
	}
}
