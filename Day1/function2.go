package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(displayName("saket", "mishra"))
}
func displayName(firstName, lastName string) (string, int) {
	fullName := firstName + " " + lastName
	return strings.ToUpper(fullName), len(fullName)
}
