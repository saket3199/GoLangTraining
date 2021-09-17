package main

import "fmt"

func main() {
	func() {

		fmt.Println("Inside Anynomous function")
	}()
	outSideMain()
}

func outSideMain() {
	func() {
		fmt.Println("Inside Anynomous function Outside Main")
	}()
}
