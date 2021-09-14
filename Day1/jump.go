package main

import "fmt"

func main() {
	var a int
start:
	fmt.Println("Enter a Number")
	fmt.Scanln(&a)
	fmt.Println("Entered Number is", a)
	if a > 100 {
		fmt.Println("Entered very high Number")
		goto start
	} else if a < 50 {

		fmt.Println("Good Number")
	} else {
		fmt.Println("Exiting!!")
	}

}
