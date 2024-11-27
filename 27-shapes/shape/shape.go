package myshape

import "fmt"

var global int

func Greet() {
	//fmt.Println("hello World")
	anotherGreet()
}

func anotherGreet() { // simialr private, restricted
	fmt.Println("Hello ICICI Securities folks")
}
