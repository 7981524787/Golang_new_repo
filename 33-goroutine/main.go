package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main is started")
	go greet()
	go func() {
		r := add(10, 12)
		fmt.Println(r)
	}()

	fmt.Println("main has exited")
	//time.Sleep(time.Millisecond * 1)

	runtime.Goexit() // it exit your main
}
func greet() {
	time.Sleep(time.Second * 2)
	c := 1
	for {

		if c >= 10 {
			runtime.Goexit()
		}
		fmt.Println("Hello ICICI")
		c++
	}
}

func add(a, b int) int {
	return a + b
}

// 1. main is also a goroutine (runtime.main.main)
// 2. No goroutine waits for other goroutine to completes its execution
// 3. order of execution is not guaranteed
// 4. easy to create just by using go keyword
// 5. if the function returns then directly you cannot run it as a go routine
