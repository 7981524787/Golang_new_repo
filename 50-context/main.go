package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generate(time.Second * 2)
	sig := make(chan bool)
	go receiver(ch, sig)
	<-sig
}

func generate(d time.Duration) chan int {
	ch1 := make(chan int)
	timeout := timeOut(d)
	// time.After(d)
	go func() {
		i := 1
	out:
		for {
			select {
			case ch1 <- i * i:
				time.Sleep(time.Millisecond * 100)
				ch1 <- i // blocked
				i++
			case <-timeout:
				close(ch1)
				break out
			}
		}
	}()
	return ch1
}

func timeOut(d time.Duration) chan struct{} {
	sig := make(chan struct{})
	go func() {
		time.Sleep(d) // some time
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}

func receiver(ch chan int, sig chan bool) {
	for v := range ch {
		fmt.Println("receiver-1", v)
	}
	sig <- true
}
