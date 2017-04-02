package main

import "fmt"

//
// celsius to fahrenheit conversion
// C = (F - 32) * 5/9
//
func main() {
	fmt.Println("Celcius to Fahrenheit Conversion")

	fmt.Print("Enter a temperature (Celcius): ")
	var celsius float32
	fmt.Scanf("%f", &celsius)

	fahrenheit := (celsius*9)/5 + 32

	fmt.Println("Temperature (Fahrenheit): ", fahrenheit)
}
