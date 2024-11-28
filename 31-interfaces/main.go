package main

import (
	"fmt"
	"os"
)

func main() {
	f := NewFile("data.txt")
	_, err := f.Write([]byte("Hello world"))
	if err != nil {
		fmt.Println(err)
	}
}

type FileWriter struct {
	Filename string
}

func NewFile(fn string) *FileWriter {
	return &FileWriter{Filename: fn}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(fw.Filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return 0, &CustomError{Code: 101, Status: "File Error", Message: err.Error()}
	}
	//errors.New()
	defer f.Close() // ignore defer as of now

	return f.Write(p)
}

type CustomError struct {
	Code    int
	Status  string
	Message string
}

func (cr *CustomError) Error() string {
	// implement any code according to your requirement

	return fmt.Sprint("Code: ", cr.Code, "Status: ", cr.Status, "Message: ", cr.Message)
}
