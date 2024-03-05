package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getUserInt() []int {
	var userInput string
	var userInputConverted []int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a series of 12 integers separated by spaces: ")
	userInput, _ = reader.ReadString('\n')
	splitString := strings.Fields(userInput)
	for _, val := range splitString {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		userInputConverted = append(userInputConverted, i)
	}
	fmt.Println(userInputConverted)
	return userInputConverted
}

func splitSlice(s []int) [][]int {
	lenSlice := len(s)
	intChunksize := int(float64(lenSlice) / 4)
	groupedIntSlice := make([][]int, 0)
	for i := 0; i < 4; i++ {
		groupedIntSlice = append(groupedIntSlice, s[i*intChunksize:i*intChunksize+intChunksize])
	}
	return groupedIntSlice
}

func sortNums(s []int, c chan []int) {
	fmt.Printf("Incoming subarray: %v \n", s)
	sort.Ints(s)
	fmt.Printf("Sorted subarray: %v \n", s)
	c <- s
}

func main() {
	intSlice := getUserInt()
	groupedSlice := splitSlice(intSlice)
	returnedArrays := make([]int, 0)
	for _, g := range groupedSlice {
		c := make(chan []int)
		go sortNums(g, c)
		sortedSubArray := <-c
		returnedArrays = append(returnedArrays, sortedSubArray...)
	}
	sort.Ints(returnedArrays)
	fmt.Printf("%v \n", returnedArrays)
}
