package main

import "fmt"

func comp(num1 int, num2 int) (isEquall bool, diff int) {

	switch {
	case num1 > num2:
		isEquall = false
		diff = num1 - num2
	case num2 > num1:
		isEquall = false
		diff = num2 - num1
	default:
		isEquall = true
		diff = 0
	}
	return

}

func main() {

	fmt.Println(comp(1, 1))
}
