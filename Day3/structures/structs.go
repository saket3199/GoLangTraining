package structures

// import "fmt"

// func main() {

// 	var person1 = Person{
// 		firstName: "Saket",
// 		lastName:  "Mishra",
// 		age:       21,
// 	}
// 	fmt.Println(person1)
// 	person2 := NewPerson("abc", "xyz", 20)
// 	fmt.Println(person2)

// 	//type embedding

// 	employee := Employee{
// 		EmployeeId: 101,
// 		Person: Person{
// 			firstName: "Saket",
// 			lastName:  "Mishra",
// 			age:       22,
// 		},
// 	}
// 	fmt.Println(employee)

// 	//calling from another file
// 	Structures()
// }

func NewPerson(fName, lName string, age int) *Person {
	var testPerson = &Person{
		firstName: fName,
		lastName:  lName,
		age:       age,
	}
	return testPerson
}

// method
func (p *Person) FullName() string {
	return p.firstName + " " + p.lastName
}

type Employee struct {
	EmployeeId int
	Person
}

type Person struct {
	firstName string
	lastName  string
	age       int
}
