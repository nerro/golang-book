package main

import "fmt"

// task3
func findMax(args ...int) int{
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

func main() {
	fmt.Println(findMax(2, 1, 17, 3, 8))

	makeOdd := makeOddGenerator()
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())
	fmt.Println(makeOdd())
}
