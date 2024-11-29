package main

import (
	"fmt"
	"os"
)

func main() {

	fw := NewFile("data.txt")

	fmt.Fprintln(fw, "Hello World")
	fmt.Fprintln(os.Stdout, "Hello World")
	ew := &EmailWriter{sender: "jp@outlook.com", receiver: "somereceiver@gmail.com", email: "dummy styff"}
	fmt.Fprintln(ew, "Hello World")

}

type EmailWriter struct {
	email    string
	sender   string
	receiver string
}

func (fw *EmailWriter) Write(p []byte) (n int, err error) {
	//
	return 0, nil
}

type FileWriter struct {
	Filename string
}

func NewFile(fn string) *FileWriter {
	return &FileWriter{Filename: fn}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(fw.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close() // ignore defer as of now

	// _, err = f.Write(p)
	// if err != nil {
	// 	f.Close()
	// }
	// f.Close()
	return 0, nil
}

type IShape interface {
	Area() float32
	Perimeter() float32
}

type Rect struct {
	L    float32
	B    float32
	A, P float32
}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Area() float32 { // function
	(*r).A = (*r).L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 { // function
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) What() string {
	return "Rect"
}
