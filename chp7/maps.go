package main

import "fmt"

func main() {
	jewerly := make(map[string]float64)
	jewerly["necklace"] = 79.99
	jewerly["earrings"] = 89.99
	clothing := map[string]float64{"Shirt": 59.99, "Pants": 49.59}
	for cloth, price := range clothing {
		fmt.Println(cloth, price)
	}

}
