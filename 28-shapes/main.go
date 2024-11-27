package main

import (
	"math/rand"

	myshape "github.com/jitenpalaparthi/icici-shapes/shape"
	"github.com/jitenpalaparthi/icici-shapes/shape/rect"
	"github.com/jitenpalaparthi/icici-shapes/shape/square"
)

func main() {

	rand.Intn(100)
	r1 := rect.New(12.23, 123.23)
	a1 := r1.Area()
	p1 := r1.Perimeter()

	println("Area of rect r1:", a1)
	println("Perimeter of rect r1:", p1)
	myshape.Greet()
	//shape.anotherGreet()
	s1 := square.Square{S: 123.23}
	a2 := s1.Area()
	p2 := s1.Perimeter()

	println("Area of square s1:", a2)
	println("Perimeter of square r1:", p2)
}
