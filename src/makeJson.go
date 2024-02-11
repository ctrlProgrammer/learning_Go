package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var userData = make(map[string]string)
	var userName string
	var userAddress string

	fmt.Println("Enter your Name: ")
	fmt.Scan(&userName)
	fmt.Println("Enter your address: ")
	fmt.Scan(&userAddress)

	userData["name"] = userName
	userData["address"] = userAddress

	jsonData, err := json.Marshal(userData)

	if err != nil {
		fmt.Println("Error formating the json.")
	} else {
		fmt.Println(string(jsonData))
	}
}
