package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a Input")
	// input := bufio.NewReader(os.Stdin)
	// userInput, _ := input.ReadString('\n')
	var userInput string
	fmt.Scan(&userInput)
	fmt.Println("User Input is: ", userInput)

}
