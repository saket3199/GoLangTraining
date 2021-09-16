package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice)
	fmt.Println(sum(mySlice...))
}
func sum(num ...int) int {
	total := 0
	for _, val := range num {
		total += val
	}
	return total
}
