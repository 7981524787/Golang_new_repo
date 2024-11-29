package main

import "fmt"

func main() {
	// defer fmt.Println("End of main-1")
	// defer fmt.Println("End of main-2")
	// generate()
	// fmt.Println("Start of main")
	fmt.Println("-----------------")
	num := 100
	defer func(v int) {
		v++
		println("with in defer", v)
	}(num)

	num = 200
	fmt.Println("num:", num)
	//

	str := "Heloo World"

	for _, v := range str {
		defer println(string(v))
	}

}

func generate() {
	defer func() {
		fmt.Println("start of generate function")
		defer fmt.Println("end of generate function")
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				println(i)
			}
		}
	}()

	fmt.Println("even number generator")
}

// defer, panic ,recover
