package main

import "fmt"

var outside = "Hello, I'm from outside"

func main() {
	var x string
	x = "Hello World"
	fmt.Println(x)

	y := "Hello World"
	fmt.Println(y)

	fmt.Println(outside)

	const z string = "constant"
	fmt.Println(z)

	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// input
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
