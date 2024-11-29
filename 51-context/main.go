package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	parent := context.Background()
	//somefunc(context.TODO(), 100)
	sigContext, cancel := signal.NotifyContext(parent, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	// //sigContext.Done()
	defer cancel()
	///	sigContext, cancel := context.WithDeadline(parent, time.Now().Add(time.Second*2)
	//defer cancel()
	// sigContext, cancel := context.WithCancel(parent)
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	cancel()
	// }()

	ch, sig := generate(sigContext)
	go func() {
		for v := range ch {
			println(v)
		}
	}()
	<-sig
}

func generate(ctx context.Context) (chan int, chan struct{}) {
	ch := make(chan int)
	sig := make(chan struct{})
	go func() {
		i := 1
	out:
		for {
			time.Sleep(time.Millisecond * 300)
			select {
			case ch <- i * i:
			case <-ctx.Done():
				time.Sleep(time.Second * 1)
				println("context cancelled in generate")
				sig <- struct{}{}
				close(ch)
				close(sig)
				break out
			}
			i++
		}
	}()
	return ch, sig
}

func somefunc(ctx context.Context, num int) {
	fmt.Println(num)
	//
}
