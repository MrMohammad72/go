package main

import "fmt"

func comp(num1 int, num2 int) (isEquall bool, diff int) {
	if num1 > num2 {
		//return false, num1 - num2
		isEquall = false
		diff = num1 - num2
	} else if num2 > num1 {
		//return false, num2 - num1
		isEquall = false
		diff = num2 - num1
	} else {
		//return true, 0
		isEquall = true
		diff = 0
	}
	return

}

func main() {

	fmt.Println(comp(1, 1))
}
