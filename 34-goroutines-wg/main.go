package main

import (
	"fmt"
	"os"
	"sync"
)

// waitgroup
// counter

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {
	// wg.Add(1) // 1
	// wg.Done() // 0
	// wg.Add(2) // 2
	// wg.Done() // 1
	// wg.Done() // 0

	wg.Add(1) // 1
	go func() {
		for i := 0; i <= 100; i++ {
			wg.Add(1) //101
			go evenNumbers(i)
		}
		wg.Done() // 100
	}()

	wg.Add(1)
	go func() {
		for i := 0; i <= 100; i++ {
			wg.Add(1)
			go oddNumbers(i)
		}
		wg.Done()
	}()

	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			os.WriteFile("data.txt", []byte(fmt.Sprint(i)), 0644)
			wg.Done()
		}(wg)
	}

	evenNumbers(100)
	//wg.Add(5)
	//wg.Wait() //
	wg.Wait() // wait until the counter inside the waitgroup becomes zero
}

func evenNumbers(n int) {
	if n%2 == 0 {
		fmt.Println("Even number->", n)
	}
	wg.Done() // -1
}

func oddNumbers(n int) {
	if n%2 != 0 {
		fmt.Println("Odd number->", n)
	}
	wg.Done()
}
