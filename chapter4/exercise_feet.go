package main

import "fmt"

//
// feet to meter conversion
// 1 feet = 0.3048 meter
//
func main() {
	fmt.Println("Feet to Meter Conversion")

	fmt.Print("Enter a length (feet): ")
	var feet float32
	fmt.Scanf("%f", &feet)

	meter := feet * 0.3048

	fmt.Println("Length (meter): ", meter)
}
