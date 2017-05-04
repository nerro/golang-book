package main

import "fmt"

func zeroCopyByValue(x int) {
	x = 0
}

func zeroCopyBeReference(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	zeroCopyByValue(x)
	fmt.Println(x)

	zeroCopyBeReference(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
