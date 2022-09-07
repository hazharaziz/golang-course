package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}

type circle struct {
	radius float64
}

type shape interface {
	calculateArea() float64
}

func (s square) calculateArea() float64 {
	return s.sideLength * s.sideLength
}

func (c circle) calculateArea() float64 {
	return math.Pi * (c.radius * c.radius)
}

func info(s shape) {
	fmt.Println(s.calculateArea())
}

func main() {
	s := square{
		sideLength: 10,
	}

	c := circle{
		radius: 3,
	}

	info(s)
	info(c)
}
