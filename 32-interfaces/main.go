package main

import "fmt"

func main() {
	c := New()
	v := c.Add(10).Add(20).Add(-20).Get()
	fmt.Println(v)
}

type Calc struct {
	Data int
}

func (c *Calc) Add(v int) ICalc {
	c.Data += v
	return c
}

func (c *Calc) Get() int {
	return c.Data
}
func New() *Calc {
	return &Calc{Data: 0}
}

type ICalc interface {
	Add(int) ICalc
	Get() int
}
