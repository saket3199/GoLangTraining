package main

import "fmt"

func main() {

	var person1 = Person{
		firstName: "Saket",
		lastName:  "Mishra",
		age:       21,
	}
	fmt.Println(person1)
	person2 := newPerson("abc", "xyz", 20)
	fmt.Println(person2)
}
func newPerson(fName, lName string, age int) *Person {
	var testPerson = &Person{
		firstName: fName,
		lastName:  lName,
		age:       age,
	}
	return testPerson
}

type Person struct {
	firstName string
	lastName  string
	age       int
}
