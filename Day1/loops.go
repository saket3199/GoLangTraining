package main

import "fmt"

func main() {
	fmt.Println("Enter a Number")
	var a int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {

		fmt.Println("Incrementing", i)
	}
	for i := a; i > 0; i-- {
		fmt.Println("Decrementing", i)
	}

}
