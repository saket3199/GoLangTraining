package main

import "fmt"

func main() {
	fmt.Println(addOne(9))
	// addOne(9)
	fmt.Println(doMathOperations(addNumber))
	fmt.Println(doMathOperations(subtractNumber))
	fmt.Println(doMathOperations(multiplyNumber))
}
func addOne(number int) int {
	number += 1
	return number
}
func addNumber(num1, num2 int) int {
	return num1 + num2
}
func multiplyNumber(num1, num2 int) int {
	return num1 * num2
}
func subtractNumber(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}
func doMathOperations(funcName func(num1, num2 int) int) int {
	result := funcName(20, 30)
	return result
}
