package main

import (
	"fmt"
)

func sum(xs []int) int {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return sum
}

func half(x int) (int, bool) {
	half := float64(x) / 2
	even := false

	//special cases
	if x == 1 {
		even = false
	} else if x == 2 {
		even = true
	} else if int(half)%2 == 0 {
		even = true
	}

	return int(half), even
}

func main() {
	//task 1
	xs := []int{1, 3, 5, 7}
	fmt.Println(sum(xs))

	//task2
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(3))
	fmt.Println(half(4))
}
