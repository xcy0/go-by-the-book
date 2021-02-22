package main

import "fmt"

func myFunc(arg ...int) {
	for _, i := range arg {
		fmt.Println(i)
	}
}

func main() {
	myFunc(1, 2, 3, 4)
}
