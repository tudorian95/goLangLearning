package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s circle) area() float64 {
	return s.radius * s.radius * math.Pi
}

func main() {
	s1 := square{4}
	c1 := circle{4}
	i1 := info(s1)
	i2 := info(c1)
	fmt.Println(i1, i2)
}
