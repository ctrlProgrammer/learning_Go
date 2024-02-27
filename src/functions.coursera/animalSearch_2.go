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

func CreateAnimal(animalType string) Animal {
	switch animalType {
	case "cow":
		return Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		return Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		return Animal{food: "mice", locomotion: "slither", noise: "hsss "}
	default:
		return Animal{food: "mice", locomotion: "slither", noise: "hsss "}
	}
}

func main() {
	var userCommand string
	var searchedAnimal string
	var searchedInfo string

	var baseAnimals = make(map[string]Animal)
	var userAnimals = make(map[string]Animal)

	baseAnimals["cow"] = CreateAnimal("cow")
	baseAnimals["bird"] = CreateAnimal("bird")
	baseAnimals["snake"] = CreateAnimal("snake")

	for {
		fmt.Println("Welcome to the ranch.")
		fmt.Println("We have some animals in the ranch, you can search the animals information using this application.")
		fmt.Println("Our main animals are the snake, the bird and the cow, please search one of these animals.")
		fmt.Println("If you want to add a new animal enter 'newanimal' or if you want to search something about the animals enter 'query' ")
		fmt.Println("If you use the 'newanimal' command you should add 'animalName' and the 'animalType' as parameters like in the next command <newanimal papo bird>")
		fmt.Println("If you use the 'query' command you should add 'animalName' and the 'animalAction', the animal can eat, move or speak, as parameters like in the next command <query papo speak>")
		fmt.Print(">")
		fmt.Scan(&userCommand, &searchedAnimal, &searchedInfo)

		switch userCommand {
		case "newanimal":
			animal, ok := baseAnimals[searchedInfo]

			if ok {
				userAnimals[searchedAnimal] = animal
				fmt.Println("Your animal (" + searchedAnimal + " the " + searchedInfo + ") was added in your ranch.")
			}

		case "query":
			animal, ok := userAnimals[searchedAnimal]

			if ok {
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

			} else {
				fmt.Println("This is an invalid animal. We dont have it in our ranch.")
			}
		default:
			fmt.Println("Invalid comand from the user.")
		}

	}

}
