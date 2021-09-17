package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	myMap["a"] = 10
	myMap["b"] = 20
	myMap["c"] = 30

	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println("The key is :", key, " and Value is :", value)
	}
	var test, ok = myMap["a"]
	fmt.Println(test, ok)
	delete(myMap, "a")
	fmt.Println(changeMap(myMap))
	// goto test
}
func changeMap(myMap map[string]int) map[string]int {
	myMap["a"] = 100
	return myMap
}
