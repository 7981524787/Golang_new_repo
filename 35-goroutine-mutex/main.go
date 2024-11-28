package main

import "sync"

var counter int = 0
var mu *sync.Mutex = new(sync.Mutex)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(302)

	go func() {
		for i := 1; i <= 200; i++ {
			go incr(wg, mu)
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			go decr(wg, mu)
		}
		wg.Done()
	}()

	wg.Wait()
	println(counter)

}

func incr(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func decr(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter--
	mu.Unlock()
	wg.Done()
}

// do not sharem memory by communicating, do communicate by sharing memory

type Counter struct {
	C  int
	Mu *sync.Mutex
}

func New(mu *sync.Mutex, c int) *Counter {
	return &Counter{C: c, Mu: mu}
}

func (c *Counter) Incr(wg *sync.WaitGroup, mu *sync.Mutex) {
	c.Mu.Lock()
	c.C++
	c.Mu.Unlock()
	wg.Done()
}

func (c *Counter) Decr(wg *sync.WaitGroup, mu *sync.Mutex) {
	c.Mu.Lock()
	c.C--
	c.Mu.Unlock()
	wg.Done()
}
