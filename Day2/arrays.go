package main

import "fmt"

func main() {
	var array = [3]string{"a", "b", "c"}
	fmt.Println("Original Array", array)
	case1(array)
	fmt.Println("After Case 1", array)
	case2(&array)
	fmt.Println("After case 2", array)
}
func case1(array [3]string) {
	fmt.Println(array[0])
}
func case2(array *[3]string) {
	array[0] = "hey"
	fmt.Println(array[0])
}
