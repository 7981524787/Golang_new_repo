package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := generate(0, 10) // generator
	<-receiver(ch1, ch2)        // future

}
func generate(f, l int) (chan int, chan int) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := f; i <= l; i++ {
			time.Sleep(time.Millisecond * 100)
			ch1 <- i // blocked
			ch2 <- i
		}
		close(ch1) // closing channel is a pattern
		close(ch2) // closing channel is a pattern
	}()
	return ch1, ch2
}
func receiver(ch1 chan int, ch2 chan int) chan struct{} {
	sig := make(chan struct{})
	go func() {
		done1, done2 := false, false
		for {
			if done1 && done2 {
				sig <- struct{}{}
				close(sig)
				break
			}
			select {
			case v1, ok1 := <-ch1:
				if ok1 {
					fmt.Println("Square of v-->", v1*v1)
				} else {
					done1 = true
				}
			case v1, ok2 := <-ch2:
				if ok2 {
					fmt.Println("Cube of v-->", v1*v1*v1)
				} else {
					done2 = true
				}
			}
		}

	}()

	return sig
}
