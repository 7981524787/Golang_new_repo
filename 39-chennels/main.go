package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := generate(0, 10) // generator
	ch2 := generate(11, 20)
	<-receiver(ch1) // future
	<-receiver(ch2) // future
}
func generate(f, l int) chan int {
	ch1 := make(chan int)
	go func() {
		for i := f; i <= l; i++ {
			time.Sleep(time.Millisecond * 100)
			ch1 <- i // blocked
		}
		close(ch1) // closing channel is a pattern
	}()
	return ch1
}
func receiver(ch2 chan int) chan empty {
	sig := make(chan empty)
	go func() {
		for v := range ch2 {
			fmt.Println("Received from receiver1-->", v)
		}
		sig <- empty{}
		close(sig)
	}()
	return sig
}

type empty struct{}
