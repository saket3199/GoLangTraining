package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 2)
	ch2 := make(chan struct{})
	wg.Add(2)
	go writeToChannel(ch)
	go readToChannel(ch)
	wg.Wait()
	close(ch)
	go func() {
		ch2 <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("Values inside range are ", i)
	}
	

}
func writeToChannel(ch chan int) {
	time.Sleep(time.Second)
	ch <- 10
	ch <- 80
	ch <- 8
	fmt.Print("Data Written to Channel Successfully\n")
	wg.Done()
}
func readToChannel(ch <-chan int) {
	value := <-ch
	fmt.Println("Data inside the channel is :", value)
	wg.Done()
}
