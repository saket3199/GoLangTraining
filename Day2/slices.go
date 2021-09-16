package main

import "fmt"

func main() {
	newSlice := make([]int, 5, 5)
	for i := 1; i < 6; i++ {
		newSlice[i-1] = 10 * i
	}
	fmt.Println(newSlice)
	fmt.Println("Length of the Slice is ", len(newSlice))
	fmt.Println("Capacity of the Slice is", cap(newSlice))
	for i := 0; i < len(newSlice); i++ {
		fmt.Println("Address of the ", i, "element in the slice is", &newSlice[i])
	}
	removingElementsFromSlice(newSlice)

}
func removingElementsFromSlice(mySlice []int) {
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println("After removing element from Slice", mySlice)
}
