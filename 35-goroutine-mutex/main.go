package main

import "sync"

var counter int = 0

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(302)

	go func() {
		for i := 1; i <= 200; i++ {
			go incr(wg)
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			go decr(wg)
		}
		wg.Done()
	}()
	println(counter)
	wg.Wait()
}

func incr(wg *sync.WaitGroup) {
	counter++
	wg.Done()
}

func decr(wg *sync.WaitGroup) {
	counter--
	wg.Done()
}

// do not sharem memory by communicating, do communicate by sharing memory
