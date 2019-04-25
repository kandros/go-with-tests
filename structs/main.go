package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return (c.radius * c.radius) * math.Pi
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return (t.base * t.height) / 2
}

func perimeter(r rectangle) float64 {
	return (r.width + r.height) * 2
}

func main() {
	fmt.Println("hello")
}
