package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	sig1 := make(chan struct{})
	sig2 := make(chan empty)
	go generate(100, ch1)
	go receiver1(ch1, ch2, sig1)
	go receiver2(ch2, sig2)

	<-sig1
	<-sig2
}

func generate(num int, ch1 chan int) {
	for i := 1; i <= num; i++ {
		ch1 <- i
	}
	close(ch1) // closing channel is a pattern
}
func receiver1(ch1 chan int, ch2 chan int, sig chan struct{}) {
	for v := range ch1 {
		fmt.Println("Received from Generator-->", v)
		ch2 <- v * v
	}
	close(ch2)
	sig <- struct{}{}
	close(sig)
}

func receiver2(ch2 chan int, sig chan empty) {
	for v := range ch2 {
		fmt.Println("Received from receiver1-->", v)
	}
	sig <- empty{}
	close(sig)
}

type empty struct{}
