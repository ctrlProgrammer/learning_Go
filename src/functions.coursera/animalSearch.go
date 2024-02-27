package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println("My favorite food is:", a.food)
}

func (a *Animal) Move() {
	fmt.Println("My locomotion method is:", a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println("To call my friends I make the next sound:", a.noise)
}

func DefineAnimals(animalsMap *map[string]Animal) {

}

func main() {
	var animals = make(map[string]Animal)

	animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss "}

	for {
		var searchedAnimal string
		var searchedInfo string

		fmt.Println("Welcome to the ranch.")
		fmt.Println("We have some animals in the ranch, you can search the animals inbformation using this application")
		fmt.Println("Our main animals are the snake, the bird and the cow, please search one of these")
		fmt.Printf("What animal you want to search: ")
		fmt.Scan(&searchedAnimal)

		animal, ok := animals[searchedAnimal]

		if ok {
			for {
				fmt.Println("Now what kind of information you want to know about the" + searchedAnimal)
				fmt.Println("Enter eat, move or speak to know more about the" + searchedAnimal)
				fmt.Print(">")
				fmt.Scan(&searchedInfo)

				switch searchedInfo {
				case "eat":
					animal.Eat()
					break
				case "move":
					animal.Move()
					break
				case "speak":
					animal.Speak()
					break
				default:
					fmt.Println("It is an invalid information")
				}

			}
		} else {
			fmt.Println("This is an invalid animal. We dont have it in our ranch.")
		}

	}

}
