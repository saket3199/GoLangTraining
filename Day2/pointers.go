package main

import "fmt"

func main() {
	a := 10
	fmt.Println("Value of a is", a)
	p := &a
	fmt.Println("Address of a is", p)
	//derefrencing
	fmt.Println("Derefrence Value of a", *p)
	//double refrencing
	q := &p
	fmt.Println("Address of p is", q)
	fmt.Println("Double de refrencing", **q)

	//pass by Value
	case1(a)
	//pass by refrence
	case2(&a)
}
func case2(number *int) {
	fmt.Println("Value of number is", *number)
	fmt.Println("Address of number is ", &number)
}
func case1(number int) {
	fmt.Println("Value of number is", number)
	fmt.Println("Address of number is ", &number)
}
