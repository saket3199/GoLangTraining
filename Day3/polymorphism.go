package main

import "fmt"

func main() {
	passInterface(&circle{
		radius: 10,
	})
	passInterface(&square{
		side: 20,
	})
}

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c *circle) area() float64 {
	return 22 / 7 * (c.radius * c.radius)
}

func (c *circle) perimeter() float64 {
	return 2 * 22 / 7 * c.radius
}

func (s *square) area() float64 {
	return s.side * s.side
}

func (s *square) perimeter() float64 {
	return 4 * s.side
}

func passInterface(sh shape) {
	fmt.Println("Area is : ", sh.area())
	fmt.Println("perimeter is : ", sh.perimeter())

}
