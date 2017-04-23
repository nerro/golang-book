package main

import "fmt"

// task3
func findMax(args ...int) int {
	max := args[0]

	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

// task4 makeOddGenerator
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(x uint) uint {
	// seed values cases
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Println("task 3")
	fmt.Println(findMax(2, 1, 17, 3, 8))

	fmt.Println("task 4")
	makeOdd := makeOddGenerator()
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())

	fmt.Println("task 5")
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(9))
}
