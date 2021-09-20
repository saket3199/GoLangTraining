package main

import (
	"chp10/geo"
	"fmt"
)

func main() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(34.5)
	coordinates.SetLongitude(56.7)
	fmt.Println(coordinates)
}
