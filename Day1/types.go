package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("Abc"))
	fmt.Println(reflect.TypeOf(4))
	fmt.Println(reflect.TypeOf(4.4))
	fmt.Println(reflect.TypeOf(true))
}
