package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hiss")
}

func main() {
	animalsMap := make(map[string]Animal)
	methodsMap := make(map[string]string)
	AddMethodsToMap(&methodsMap)

	for {
		action, token1, token2, err := GetInputFromUser()
		if err != nil {
			fmt.Println(err)
		} else {
			isNewAnimal := strings.Compare(action, "newanimal") == 0
			isQuery := strings.Compare(action, "query") == 0
			if !isNewAnimal && !isQuery {
				fmt.Println("Action must be either 'newanimal' or 'query'")
			} else if isNewAnimal {
				createNewAnimal(token1, token2, &animalsMap)
			} else {
				query(token1, token2, &animalsMap, &methodsMap)
			}
		}
	}
}

func AddMethodsToMap(methodsMapPtr *map[string]string) {
	methodsMap := *methodsMapPtr
	methodsMap["eat"] = "Eat"
	methodsMap["move"] = "Move"
	methodsMap["speak"] = "Speak"
}

func GetInputFromUser() (string, string, string, error) {
	in := bufio.NewReader(os.Stdin)

	fmt.Print(">")
	str, err := in.ReadString('\n')
	if err != nil {
		return "", "", "", err
	}
	str = strings.TrimSuffix(str, "\n")
	tokens := strings.Split(str, " ")

	if len(tokens) != 3 {
		return "", "", "", errors.New("Incorrect number of args.")
	}
	return tokens[0], tokens[1], tokens[2], nil
}

func createNewAnimal(name, animalType string, animalsMapPtr *map[string]Animal) {
	animalsMap := *animalsMapPtr
	if strings.Compare(animalType, "cow") == 0 {
		animalsMap[name] = Cow{}
		fmt.Println("Created it!")
	} else if strings.Compare(animalType, "bird") == 0 {
		animalsMap[name] = Bird{}
		fmt.Println("Created it!")
	} else if strings.Compare(animalType, "snake") == 0 {
		animalsMap[name] = Snake{}
		fmt.Println("Created it!")
	} else {
		fmt.Printf("Invalid animal type: %s\n", animalType)
	}
}

func query(name, method string, animalsMapPtr *map[string]Animal, methodsMapPtr *map[string]string) {
	animalsMap := *animalsMapPtr
	methodsMap := *methodsMapPtr
	if _, ok := animalsMap[name]; !ok {
		fmt.Printf("Animal name not found: %s\n", name)
	} else if _, ok := methodsMap[method]; !ok {
		fmt.Printf("Method not found: %s\n", method)
	} else {
		reflect.ValueOf(animalsMap[name]).MethodByName(methodsMap[method]).Call([]reflect.Value{})
	}
}
