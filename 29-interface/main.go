package main

import "fmt"

func main() {

	r1 := New(123.123, 123.23)
	var s1 Square = 123.123
	// Shape(r1)
	// Shape(s1)

	slice := make([]IShape, 2)
	slice[0] = r1
	slice[1] = s1

	for _, v := range slice {
		Shape(v)
	}
	// Cuboid
	// creat a slice of interfaces
	// all all object like 2 rect , 2 square , 1 cuboid and run them in a loop
}

func Shape(ishape IShape) {
	fmt.Println("Area of:", ishape.What(), "is:", ishape.Area())
	fmt.Println("Perimeter of:", ishape.What(), "is:", ishape.Perimeter())
}

type IShape interface {
	// Area() float32
	// Perimeter() float32
	IArea
	IPerimeter
	IWhat
}

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IWhat interface {
	What() string
}

type Rect struct {
	L    float32
	B    float32
	A, P float32
}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Area() float32 { // function
	(*r).A = (*r).L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 { // function
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) What() string {
	return "Rect"
}

type Square float32

func (s Square) Area() float32 { // function
	return float32(s * s)
}

func (s Square) Perimeter() float32 { // function
	return float32(4 * s)
}

func (s Square) What() string {
	return "Square"
}
