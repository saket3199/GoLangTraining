package main

import "fmt"

func main() {
	var pet struct {
		name string
		age  int
	}
	pet.name = "Max"
	pet.age = 2
	fmt.Println("Name :", pet.name)
	fmt.Println("age :", pet.age, " years")
}
