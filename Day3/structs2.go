package main

import (
	"fmt"
	"structures"
)

func main() {

	person3 := structures.NewPerson("abc", "xyz", 21)
	fmt.Println(*person3)
	fmt.Println(person3.FullName())
}
