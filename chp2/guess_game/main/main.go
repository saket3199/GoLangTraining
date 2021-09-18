package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	target := rand.Int63n(100) + 1
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println("Enter a Guess")
		input, _ := reader.ReadString('\n')
		guess, _ := strconv.ParseInt(input, 0, 32)
		fmt.Println(guess)
		if guess == target {
			fmt.Println("Yay!! You guessed it right !")
			break
		} else if guess < target {
			fmt.Println("Oops Your Guess is Low !")
			fmt.Println("Guesses left :", 9-i)
		} else {
			fmt.Println("Oops Your Guess is High !")
			fmt.Println("Guesses left :", 9-i)
		}

	}

}
