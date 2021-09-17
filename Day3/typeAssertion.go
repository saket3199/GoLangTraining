package main

import "fmt"

func main() {
	var i interface{} = 10
	storeString, ok := i.(int)
	fmt.Println(storeString, ok)
	var num float32 = 17
	checkType(num)
	checkType(15)
	checkType(true)
}

func checkType(anyType interface{}) {
	switch anyType.(type) {
	case string:
		fmt.Println("Argument is String")
	case float64:
		fmt.Println("Argument is float64")
	case float32:
		fmt.Println("Argument is float32")
	case int:
		fmt.Println("Argument is int")
	default:
		fmt.Println("other than my cases")
	}
}
