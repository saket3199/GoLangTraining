package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	// var wg = sync.WaitGroup{}
	wg.Add(2)
	go delayedTimeFunction()
	go delayedTimeFunction2()
	wg.Wait()

}
func delayedTimeFunction() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Inside method 1 ", i)
	}
	wg.Done()
}

func delayedTimeFunction2() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Inside method 2 ", i)
	}
	wg.Done()
}
