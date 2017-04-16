package main

import "fmt"

func main() {
	second := make([]int, 3, 9)
	fmt.Println(len(second))

	third := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(third[2:5])

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	minValue := x[0]
	for _, value := range x {
		if minValue > value {
			minValue = value
		}
	}
	fmt.Println("Minimal value is:", minValue)
}
