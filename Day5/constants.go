package main

import "fmt"

const myConstant int = 30

const (
	_ = iota
	cat
	dog
	horse
	cow
)

func main() {
	fmt.Println("Value of constant is", myConstant)
	fmt.Println(cat, dog, horse, cow)
}
