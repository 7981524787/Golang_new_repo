package main

import (
	"fmt"
	"runtime"
)

func main() {

	var ch chan int // declared not instantiated

	ch = make(chan int) // unbuffered channel instantiated

	go func() {
		v := <-ch //// sender , this is blocked until the receiver receives the value
		println(v)
	}()

	ch <- 100

	// buffered channel

	ch2 := make(chan int, 2)
	ch2 <- 200
	ch2 <- 300
	ch2 <- 400
	v2 := <-ch2

	fmt.Println(v2)

	runtime.Goexit()
}

//
