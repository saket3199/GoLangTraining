package main

import "fmt"

func main() {

	f2()
	defer f1()
}
func f1() {
	fmt.Println("Inside F1")
}
func f2() {
	defer fmt.Println("Inside F1 defer")
	fmt.Println("Inside F2")
}
