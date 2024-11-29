package main

import "fmt"

func main() {
	//var num = 0
	//println(100 / num)

	fmt.Println("Hello World")
	func() {
		defer fmt.Println("calling divide function")
		defer func() {
			for i := 10; i >= 1; i-- {
				println("numbers before defer", i)
			}
		}()
		divide(20)
	}()

	for i := 100; i >= 1; i-- {
		println("numbers", i)
	}

}

func divide(num int) {
	for i := num; i >= 0; i-- {
		println(num / i)
	}
}

// defer, panic ,recover
