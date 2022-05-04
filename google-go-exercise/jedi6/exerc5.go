package main

import (
	"fmt"
	"math"
)

// type Shape interface {
// 	CalculateArea() float32
// }

type Square struct {
	base   float32
	height float32
}

type Circle struct {
	r float32
}

func (s Square) CalculateArea() float32 {
	return s.base * s.height
}

func (c Circle) CalculateArea() float32 {
	return math.Pi + c.r*2
}

func main() {
	square := Square{base: 10, height: 10}
	circle := Circle{r: 15}

	fmt.Println(square.CalculateArea())
	fmt.Println(circle.CalculateArea())
}
