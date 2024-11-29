package main

import "fmt"

func main() {

	Fullname("J", "P")
	func() {
		defer recoverthis()
		Fullname("J", "")
		for i := 1; i <= 10; i++ {
			fmt.Println("inside", i)
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println("outside", i)
	}

}

func Fullname(fn, ln string) {

	if fn == "" || ln == "" {
		panic("fistname or lastname is empty")
	}
	fmt.Println(fn, ln)
}

func recoverthis() {
	if r := recover(); r != nil {
		fmt.Println("recovering from this panic", r)
	}
}
