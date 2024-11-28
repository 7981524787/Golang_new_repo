package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	sig1 := make(chan struct{})
	sig2 := make(chan empty)
	go generate(100, ch)
	go receiver1(ch, sig1)
	go receiver2(ch, sig2)
	// v := <-sig
	// fmt.Println(v)
	<-sig1
	<-sig2
}

func generate(num int, ch chan int) {
	for i := 1; i <= num; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- i * i
	}
	close(ch) // closing channel is a pattern
}
func receiver1(ch chan int, sig chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			sig <- struct{}{}
			runtime.Goexit()
		} else {
			fmt.Println("receiver-1", v)
		}
	}
}

func receiver2(ch chan int, sig chan empty) {
	for v := range ch {
		fmt.Println("receiver-2", v)
	}
	sig <- empty{}
}

type empty struct{}
