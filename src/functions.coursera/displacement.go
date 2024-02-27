package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a float64, iv float64, id float64) func(t float64) float64 {
	return func(t float64) float64 {
		return id + iv*t + (a*(t*t))/2
	}
}

func main() {
	//Acceleration, Initial velocity and initial displacement
	var a, iv, id float64

	fmt.Println("Welcome to the displacement calculator: ")
	fmt.Println("Enter the acceleration, the initial velocity and the initial displacement to calculate the total displacement over the time.")

	fmt.Printf("Enter the aceleration:")
	fmt.Scan(&a)
	fmt.Printf("Enter the initial velocity:")
	fmt.Scan(&iv)
	fmt.Printf("Enter the initial displacement:")
	fmt.Scan(&id)

	fn := GenDisplaceFn(a, iv, id)

	for i := 0; i < 10; i++ {
		fmt.Println("Displacement before "+strconv.Itoa(i*3)+" seconds: ", fn(float64(i*3)))
	}
}
